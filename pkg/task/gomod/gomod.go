package gomod

import (
	"errors"

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
	return "dirs"
}

// Do defines action of the task
func (d *Dirs) Do() error {
	if err := exec.Run("go", "mod", "init"); err != nil {
		return err
	}
	return exec.Run("go", "mod", "tidy")
}
