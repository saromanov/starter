package project

import (
	"errors"
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/exec"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
	"github.com/saromanov/starter/pkg/task/dockerfile"
	"github.com/saromanov/starter/pkg/task/makefile"
)

var errNoName = errors.New("name of the project is not defined")

// Build provides building of the project
func Build(p *models.Project) error {
	tasks := []task.Task{}
	if p.Dockerfile != "" {
		tasks = append(tasks, dockerfile.New(p))
	}
	if p.Makefile != "" {
		tasks = append(tasks, makefile.New(p))
	}
	if err := makeDirs(p); err != nil {
		return err
	}
	if err := initGoMod(); err != nil {
		return err
	}
	return nil
}

// initGoMod provides initialization of go modules
func initGoMod() error {
	if err := exec.Run("go", "mod", "init"); err != nil {
		return err
	}
	return exec.Run("go", "mod", "tidy")
}

// makeDirs provides creating of dir structure of the project
func makeDirs(p *models.Project) error {
	if p.Name == "" {
		return errNoName
	}
	if err := os.Mkdir(p.Name, os.ModeDir); err != nil {
		return fmt.Errorf("unable to create root dir: %v", err)
	}

	for _, d := range p.SubDirs {
		if err := os.Mkdir(p.Name, os.ModeDir); err != nil {
			return fmt.Errorf("unable to create dir %s dir: %v", d, err)
		}
	}

	return nil
}
