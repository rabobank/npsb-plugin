module github.com/rabobank/npsb-plugin

go 1.23

replace (
	golang.org/x/text => golang.org/x/text v0.14.0
	google.golang.org/grpc => google.golang.org/grpc v1.62.1
	google.golang.org/protobuf => google.golang.org/protobuf v1.33.0
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
)

require code.cloudfoundry.org/cli v7.1.0+incompatible

require (
	code.cloudfoundry.org/bytefmt v0.10.0 // indirect
	code.cloudfoundry.org/tlsconfig v0.5.0 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/charlievieth/fs v0.0.3 // indirect
	github.com/cloudfoundry/bosh-cli v6.4.1+incompatible // indirect
	github.com/cloudfoundry/bosh-utils v0.0.496 // indirect
	github.com/cppforlife/go-patch v0.2.0 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/jessevdk/go-flags v1.6.1 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/moby/moby v27.3.1+incompatible // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/vito/go-interact v1.0.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/term v0.24.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
