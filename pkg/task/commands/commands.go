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

func generate() string {
	s := `
func Call(args []string) {
	app := &cli.App{
		Name:  %s,
		Commands: %s
	}
	`
	return s
}

func generateCommand(name string) string {
	return `
	{
		Name:  "%s",
		Action: func(c *cli.Context) error {
			return nil
		},
	},
	`
}
