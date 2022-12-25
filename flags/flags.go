package flags

import "github.com/urfave/cli/v2"

var (
	LogFile = &cli.StringFlag{
		Name:  "log-file",
		Value: "debug.log",
	}
)
