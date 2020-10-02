package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/mkideal/cli"
)

var versionRegexp = regexp.MustCompile(`^\d{1,2}\.\d{1,2}\.\d{1,2}$`)

type argT struct {
	Help bool `cli:"h,help" usage:"show help"`

	Name        string `cli:"n,name"       usage:"your application name"`
	Prefix      string `cli:"p,prefix"     usage:"your repository prefix (ex github.com/YusukeKishino)"`
	GoVersion   string `cli:"g,go-version" usage:"go version for goenv (ex 1.15.2)"`
	GoMod       bool   `cli:"go-mod"       usage:"use Go Module (default: true)" dft:"true"`
	Out         string `cli:"o,out"        usage:"out put directory (default: ./)" dft:"./"`
	NodeVersion string `cli:"node-version" usage:"node version for nodenv (ex 14.5.0)"`
}

func (argv *argT) AutoHelp() bool {
	return argv.Help
}

func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.Name == "" {
		return fmt.Errorf("required parameter --name missing")
	}
	if argv.Prefix == "" {
		return fmt.Errorf("required parameter --prefix missing")
	}
	if argv.GoVersion == "" {
		return fmt.Errorf("required parameter --go-version missing")
	}
	if argv.NodeVersion == "" {
		return fmt.Errorf("required parameter --node-version missing")
	}
	if !versionRegexp.MatchString(argv.GoVersion) {
		return fmt.Errorf("invalid go version %s", argv.GoVersion)
	}
	if !versionRegexp.MatchString(argv.NodeVersion) {
		return fmt.Errorf("invalid node version %s", argv.NodeVersion)
	}

	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		arg := ctx.Argv().(*argT)

		if err := generate(arg); err != nil {
			return err
		}

		return nil
	}))
}
