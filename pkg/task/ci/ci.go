// Package ci provides manipulation with ci/cd on remote services
package ci

import (
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// Task defines task for makign directories
type Task struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Task{
		p: p,
	}
}

// String returns representation of the task
func (d *Task) String() string {
	return "ci"
}

// Do defines action of the task
func (d *Task) Do() error {

	if len(d.p.CI) == 0 {
		return nil
	}

	for _, a := range d.p.CI {
		if a == models.Github {

		}
	}

	return nil
}
