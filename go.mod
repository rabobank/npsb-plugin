module github.com/rabobank/npsb-plugin

go 1.23.1

toolchain go1.23.2

replace (
	code.cloudfoundry.org/cli/v8 => github.com/jcvrabo/cli/v8 v8.8.2-test
	golang.org/x/sys => golang.org/x/sys v0.28.0
	golang.org/x/term => golang.org/x/term v0.27.0
	golang.org/x/text => golang.org/x/text v0.21.0
	google.golang.org/grpc => google.golang.org/grpc v1.69.2
	google.golang.org/protobuf => google.golang.org/protobuf v1.36.1
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
)

require code.cloudfoundry.org/cli/v8 v8.0.0-00010101000000-000000000000

require (
	code.cloudfoundry.org/bytefmt v0.24.0 // indirect
	code.cloudfoundry.org/cli v7.1.0+incompatible // indirect
	code.cloudfoundry.org/cli-plugin-repo v0.0.0-20241211062834-f5656a04bcb6 // indirect
	code.cloudfoundry.org/jsonry v1.1.4 // indirect
	code.cloudfoundry.org/tlsconfig v0.14.0 // indirect
	github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f // indirect
	github.com/charlievieth/fs v0.0.3 // indirect
	github.com/cloudfoundry/bosh-cli v6.4.1+incompatible // indirect
	github.com/cloudfoundry/bosh-utils v0.0.518 // indirect
	github.com/cppforlife/go-patch v0.2.0 // indirect
	github.com/cyphar/filepath-securejoin v0.3.6 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/jessevdk/go-flags v1.6.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/vito/go-interact v1.0.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250102185135-69823020774d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250102185135-69823020774d // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
