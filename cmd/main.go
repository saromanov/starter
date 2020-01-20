package main

import (
	"os"

	"github.com/saromanov/starter/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func build(c *cli.Context) error {
	if err := cmd.Build(); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "starter",
		Usage: "create puppet for the project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "include-dirs",
				Value: "cmd,pkg",
				Usage: "including directories to the project",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "building of the new project",
				Action: build,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
