package main

import (
	"bytes"
	"code.cloudfoundry.org/cli/cf/i18n"
	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/plugin"
	pluginmodels "code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cli/util/configv3"
	"encoding/json"
	"fmt"
	"github.com/rabobank/npsb-plugin/version"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	HttpTimeoutDefault = 5

	ListSourcesHelpText = "lists all available network policies"
	ListSourcesUsage    = "sources"
)

var (
	accessToken   string
	currentOrg    pluginmodels.Organization
	currentSpace  pluginmodels.Space
	currentUser   string
	requestHeader http.Header
	httpClient    http.Client
)

// Run must be implemented by any plugin because it is part of the plugin interface defined by the core CLI.
//
// Run(....) is the entry point when the core CLI is invoking a command defined by a plugin.
// The first parameter, plugin.CliConnection, is a struct that can be used to invoke cli commands. The second parameter, args, is a slice of strings.
// args[0] will be the name of the command, and will be followed by any additional arguments a cli user typed in.
//
// Any error handling should be handled with the plugin itself (this means printing user facing errors).
// The CLI will exit 0 if the plugin exits 0 and will exit 1 should the plugin exits nonzero.
func (c *NpsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	httpClient = http.Client{Timeout: time.Duration(HttpTimeoutDefault) * time.Second}
	if args[0] != "install-plugin" && args[0] != "CLI-MESSAGE-UNINSTALL" {
		preCheck(cliConnection)
		requestHeader = map[string][]string{"Content-Type": {"application/json"}, "Authorization": {accessToken}}
	}

	switch args[0] {
	case "sources":
		getSources(cliConnection)
	}
}

// GetMetadata returns a PluginMetadata struct. The first field, Name, determines the name of the plugin which should generally be without spaces.
// If there are spaces in the name a user will need to properly quote the name during uninstall otherwise the name will be treated as separate arguments.
// The second value is a slice of Command structs. Our slice only contains one Command Struct, but could contain any number of them.
// The first field Name defines the command `cf basic-plugin-command` once installed into the CLI.
// The second field, HelpText, is used by the core CLI to display help information to the user in the core commands `cf help`, `cf`, or `cf -h`.
func (c *NpsbPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:          "npsb",
		Version:       plugin.VersionType{Major: version.GetMajorVersion(), Minor: version.GetMinorVersion(), Build: version.GetPatchVersion()},
		MinCliVersion: plugin.VersionType{Major: 6, Minor: 7, Build: 0},
		Commands: []plugin.Command{
			{Name: "sources", HelpText: ListSourcesHelpText, UsageDetails: plugin.Usage{Usage: ListSourcesUsage}},
		},
	}
}

// preCheck Does all common validations, like being logged in, and having a targeted org and space, and if there is an instance of the scheduler-service.
func preCheck(cliConnection plugin.CliConnection) {
	config, _ := configv3.LoadConfig()
	i18n.T = i18n.Init(config)
	loggedIn, err := cliConnection.IsLoggedIn()
	if err != nil || !loggedIn {
		fmt.Println(terminal.NotLoggedInText())
		os.Exit(1)
	}
	currentUser, _ = cliConnection.Username()
	hasOrg, err := cliConnection.HasOrganization()
	if err != nil || !hasOrg {
		fmt.Println(terminal.FailureColor("please target your org/space first"))
		os.Exit(1)
	}
	org, _ := cliConnection.GetCurrentOrg()
	currentOrg = org
	hasSpace, err := cliConnection.HasSpace()
	if err != nil || !hasSpace {
		fmt.Println(terminal.FailureColor("please target your space first"))
		os.Exit(1)
	}
	space, _ := cliConnection.GetCurrentSpace()
	currentSpace = space
	if accessToken, err = cliConnection.AccessToken(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getSources(cliConnection plugin.CliConnection) {
	request := GenericRequest{SpaceGUID: currentSpace.Guid}
	requestBody, _ := json.Marshal(request)
	if cfApiEndpoint, err := cliConnection.ApiEndpoint(); err != nil {
		fmt.Printf("failed to get api endpoint: %s\n", err)
		os.Exit(1)
	} else {
		apiEndpoint := strings.Replace(cfApiEndpoint, "api.sys", "npsb.apps", 1)
		requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/sources", apiEndpoint))
		httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
		fmt.Printf("Getting source network policies for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
		resp, err := httpClient.Do(&httpRequest)
		if err != nil {
			fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from npsb service: %s", err)))
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to list source network policies: %s", err)))
			os.Exit(1)
		}
		body, _ := io.ReadAll(resp.Body)
		jsonResponse := SourceListResponse{}
		err = json.Unmarshal(body, &jsonResponse)
		if err != nil {
			fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
		}
		table := terminal.NewTable([]string{"Name", "Org", "Space", "Scope", "Description"})
		for _, src := range jsonResponse.Sources {
			table.Add(src.Source, src.Org, src.Space, src.Scope, src.Description)
		}
		_ = table.PrintTo(os.Stdout)
	}
}

// Unlike most Go programs, the `Main()` function will not be used to run all of the commands provided in your plugin.
// Main will be used to initialize the plugin process, as well as any dependencies you might require for your plugin.
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct implementing the interface defined at "code.cloudfoundry.org/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect metadata. The plugin will exit 0 and the Run([]string) method will not be invoked.
	plugin.Start(new(NpsbPlugin))
	// Plugin code should be written in the Run([]string) method, ensuring the plugin environment is bootstrapped.
}
