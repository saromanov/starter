package dirs

import (
	"errors"
	"fmt"
	"os"

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
	if d.p.Name == "" {
		return errNoName
	}
	if err := os.Mkdir(d.p.Name, os.ModeDir); err != nil {
		return fmt.Errorf("unable to create root dir: %v", err)
	}

	for _, dir := range d.p.SubDirs {
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			return fmt.Errorf("unable to create dir %s dir: %v", dir, err)
		}
	}
	return nil
}
