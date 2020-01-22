package project

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/starter/pkg/models"
)

var errNoName = errors.New("name of the project is not defined")

// Build provides building of the project
func Build(p *models.Project) error {
	if err := makeDirs(p); err != nil {
		return err
	}
	if err := createDockerfile(p); err != nil {
		return err
	}
	return nil
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

// createDockerfile provides creating of teh docker file
// if this is defined at the config
func createDockerfile(p *models.Project) error {
	if p.Dockerfile == "default" {
		return createDefaultDockerfile(p)
	}

	return moveDockerfile(p)
}

// creating of default dockerfile
func createDefaultDockerfile(p *models.Project) error {
	d1 := []byte("FROM golang:1.13")
	err := ioutil.WriteFile("Dockerfile", d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to create dockerfile: %v", err)
	}
	return nil
}

// moveDockerfile provides copy of the target Dockerfile
// to the project
func moveDockerfile(p *models.Project) error {
	return moveFile(p.Dockerfile, p.Name)
}
