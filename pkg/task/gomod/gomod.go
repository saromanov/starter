package gomod

import (
	"errors"
	"fmt"

	"github.com/saromanov/starter/pkg/exec"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

var errNoName = errors.New("name of the project is not defined")

// Dirs defines task for makign directories
type Dirs struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Dirs{
		p: p,
	}
}

// String returns representation of the task
func (d *Dirs) String() string {
	return "gomod"
}

// Do defines action of the task
func (d *Dirs) Do() error {
	if err := exec.Run(d.p.Name, "go", "mod", "init"); err != nil {
		return fmt.Errorf("unable to run 'go mod init' command: %v", err)
	}
	if err := exec.Run(d.p.Name, "go", "mod", "tidy"); err != nil {
		return fmt.Errorf("unable to run 'go mod tidy' command: %v", err)
	}
	return nil
}
