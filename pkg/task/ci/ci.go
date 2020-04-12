// Package ci provides manipulation with ci/cd on remote services
package ci

import (
	"fmt"
	"os"

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
			makeGithubActions(d.p.Name)
		}
	}

	return nil
}

// makeGithubActiosn creates subdirs for ci/cd pipeline on Github
func makeGithubActions(path string) error {
	wfPath := fmt.Sprintf("%s/%s/%s", path, ".github", "workflows")
	if err := os.MkdirAll(wfPath, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create github flow")
	}
	return task.MoveFile("./assets/go.yml", wfPath+"/go.yml")
}
