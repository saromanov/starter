package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "starter",
		Usage: "create puppet for the project",
		Commands: []*cli.Command{
			{
				Name:  "build",
				Usage: "building of the new project",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
