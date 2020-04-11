// Package ci provides manipulation with ci/cd on remote services
package ci

import (
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// CIActions defines supported type for ci/cd
type CIActions string

var (
	Github CIActions = "github"
	Travis CIActions = "travis"
	Unsupported CIActions = "unsupported"
)

// ToCIActions converts string to ci type
func ToCIActions(s string) CIActions {
	if s == "github" {
		return Github
	} else if s == "travis" {
		return Travis
	}
	return Unsupported
}

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
		if a == string(Github {

		}
	}

	return nil
}
