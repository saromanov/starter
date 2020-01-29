package project

import (
	"fmt"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
	"github.com/saromanov/starter/pkg/task/dirs"
	"github.com/saromanov/starter/pkg/task/dockerfile"
	"github.com/saromanov/starter/pkg/task/gomod"
	"github.com/saromanov/starter/pkg/task/makefile"
)

// Build provides building of the project
func Build(p *models.Project) error {
	tasks := []task.Task{dirs.New(p), gomod.New(p)}
	if p.Dockerfile != "" {
		tasks = append(tasks, dockerfile.New(p))
	}
	if p.Makefile != "" {
		tasks = append(tasks, makefile.New(p))
	}
	return runTasks(tasks)
}

// runTasks provides execution of tasks
func runTasks(tasks []task.Task) error {
	for _, t := range tasks {
		if err := t.Do(); err != nil {
			return fmt.Errorf("unable to execute task: '%s' %v", t.String(), err)
		}
	}

	return nil
}
