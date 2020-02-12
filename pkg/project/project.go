package project

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
	"github.com/saromanov/starter/pkg/task/dirs"
	"github.com/saromanov/starter/pkg/task/dockerfile"
	"github.com/saromanov/starter/pkg/task/entryfile"
	"github.com/saromanov/starter/pkg/task/gitremote"
	"github.com/saromanov/starter/pkg/task/gomod"
	"github.com/saromanov/starter/pkg/task/makefile"
	"github.com/saromanov/starter/pkg/task/readme"
)

// Build provides building of the project
func Build(p *models.Project) error {
	tasks := []task.Task{dirs.New(p), gomod.New(p), readme.New(p), gitremote.New(p)}
	if p.Dockerfile != "" {
		tasks = append(tasks, dockerfile.New(p))
	}
	if p.Makefile != "" {
		tasks = append(tasks, makefile.New(p))
	}
	switch p.Type {
	case models.Library:
		tasks = append(tasks, buildLibrary(p)...)
	case models.Binary:
		tasks = append(tasks, buildBinary(p)...)
	default:
		return errors.New("unable to define type of the project")
	}
	return runTasks(tasks)
}

// buildLibrary provides building of the library project
func buildLibrary(p *models.Project) []task.Task {
	tasks := []task.Task{}
	if p.EntryFile != "" {
		tasks = append(tasks, entryfile.New(p))
	}
	return tasks
}

// buildBinary provides bulding of the binary project
func buildBinary(p *models.Project) []task.Task {
	tasks := []task.Task{}
	return tasks
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
