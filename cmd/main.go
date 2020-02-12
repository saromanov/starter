package main

import (
	"os"

	"github.com/saromanov/starter/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func build(c *cli.Context) error {
	projectFlag := c.String("project")
	config := c.String("config")
	if err := cmd.Build(projectFlag, config); err != nil {
		logrus.Fatalf("unable to run project project: %v", err)
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
			&cli.StringFlag{
				Name:  "config",
				Usage: "path to the config",
			},
			&cli.StringFlag{
				Name:  "project",
				Value: "bin",
				Usage: "type of the project. It might be bin or lib",
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