package project

import (
	"errors"
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/models"
)

var errNoName = errors.New("name of the project is not defined")

// Build provides building of the project
func Build(p *models.Project) error {
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
