package main

import (
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func build(c *cli.Context) error {
	projectFlag := c.String("project")
	if err := cmd.Build(projectFlag); err != nil {
		logrus.Fatalf("unable to run project project: %v", err)
	}

	return nil
}

func config(c *cli.Context) error {
	conf := c.Args().First()
	if conf == "" {
		return fmt.Errorf("config is not defined")
	}
	if err := cmd.BuildFromConfig(conf); err != nil {
		logrus.Fatalf("unable to run project project: %v", err)
	}

	return nil
}

func listTemplates(c *cli.Context) error {
	if err := cmd.ListTemplates(); err != nil {
		logrus.Fatalf("unable to get list of templates: %v", err)
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
				Name:  "project",
				Value: "bin",
				Usage: "type of the project. It might be bin or lib",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "new",
				Usage:  "creating of the new project",
				Action: build,
			},
			{
				Name:   "config",
				Usage:  "building of the new project from config",
				Action: config,
			},
			{
				Name:   "list",
				Usage:  "get list of templates",
				Action: listTemplates,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
