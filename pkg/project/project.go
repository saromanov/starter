package project

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
	"github.com/saromanov/starter/pkg/task/ci"
	"github.com/saromanov/starter/pkg/task/commands"
	"github.com/saromanov/starter/pkg/task/dirs"
	"github.com/saromanov/starter/pkg/task/dockerfile"
	"github.com/saromanov/starter/pkg/task/entryfile"
	"github.com/saromanov/starter/pkg/task/gitremote"
	"github.com/saromanov/starter/pkg/task/gomod"
	"github.com/saromanov/starter/pkg/task/hosting/github"
	"github.com/saromanov/starter/pkg/task/makefile"
	"github.com/saromanov/starter/pkg/task/readme"
)

// Build provides building of the project
// first, its setting default tasks and then
// setting only non empty tasks
func Build(p *models.Project) error {
	tasks := task.Tasks{dirs.New(p), gomod.New(p), readme.New(p)}
	if p.Dockerfile != "" {
		tasks = append(tasks, dockerfile.New(p))
	}
	if p.Makefile != "" {
		tasks = append(tasks, makefile.New(p))
	}
	if p.EntryFile != "" {
		tasks = append(tasks, entryfile.New(p))
	}
	if p.HostingDescription != "" {
		tasks = append(tasks, github.New(p))
	} else {
		tasks = append(tasks, gitremote.New(p))
	}
	if len(p.CI) > 0 {
		tasks = append(tasks, ci.New(p))
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
	return tasks
}

// buildBinary provides bulding of the binary project
func buildBinary(p *models.Project) []task.Task {
	tasks := []task.Task{}
	if len(p.Commands) > 0 {
		tasks = append(tasks, commands.New(p))
	}
	return tasks
}

// runTasks provides execution of tasks
// it contains tasks on sub dir
func runTasks(tasks []task.Task) error {
	if len(tasks) == 0 {
		return nil
	}
	for _, t := range tasks {
		fmt.Printf("Starting of %s task\n", t.String())
		if err := t.Do(); err != nil {
			return fmt.Errorf("unable to execute task: '%s' %v", t.String(), err)
		}
	}

	return nil
}
