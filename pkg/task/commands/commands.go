// Package commands provides genneration for command line tools
package commands

import (
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

type Commands struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Commands{
		p: p,
	}
}

// Do provides running of the tasks
func (d *Commands) Do() error {
	if err := os.Mkdir("./pkg/cmd", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create dir %v", err)
	}
	return nil
}

func (d *Commands) String() string {
	return "commands"
}

// Check is entry point for the app
func Check(args []string) {
	app := &cli.App{
		Name:  "fresh",
		Usage: "Checking of newest deps",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "update-all",
				Usage: "updating all depencencies",
			},
			&cli.StringFlag{
				Name:  "github-token",
				Usage: "token for access to Github",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "check",
				Usage: "starting of checking",
				Action: func(c *cli.Context) error {
					if _, err := check(c, "go.mod"); err != nil {
						log.Fatalf("unable to check: %v", err)
					}
					return nil
				},
			},
			{
				Name:  "check-and-update",
				Usage: "checking of the deps and then update it",
				Action: func(c *cli.Context) error {
					if err := checkAndUpdate(c, "go.mod"); err != nil {
						log.Fatalf("unable to check and update: %v", err)
					}
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "updating of deps",
				Action: func(c *cli.Context) error {
					if err := update(c, "go.mod"); err != nil {
						log.Fatalf("unable to check: %v", err)
					}
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
func generate() string {
	s := ```
func Entry(args []string) {
	app := &cli.App{
		Name:  "%s",
		Commands: %s
	}
	```
	return s
}

func generateCommand(name string ) string {
	return ```
	{
		Name:  "%s",
		Action: func(c *cli.Context) error {
			return nil
		},
	},
	```
}
