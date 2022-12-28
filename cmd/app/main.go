package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"instagram-cli-go/flags"
	applog "instagram-cli-go/log"
	"instagram-cli-go/program"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "instagram-giveaway-cli-go",
		Usage: "CLI application for processing Instagram post details for giveaways",
		Flags: []cli.Flag{
			flags.LogFile,
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error while running application: %v",
			err)
	}
}

func run(c *cli.Context) error {
	logConfig := applog.NewConfigFromCLI(c)
	logFile, err := applog.Init(logConfig)
	if err != nil {
		return fmt.Errorf("could not initialize logger: %w", err)
	}
	defer func(logFile *os.File) {
		_ = logFile.Close()
	}(logFile)

	p := program.New()
	_, err = p.Run()
	return err
}
