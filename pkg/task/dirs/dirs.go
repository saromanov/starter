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
	rootDir := d.p.Name
	if err := os.Mkdir(rootDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create root dir: %v", err)
	}

	for _, dir := range d.p.SubDirs {
		name := fmt.Sprintf("%s/%s", rootDir, dir)
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("unable to create dir %s: %v", name, err)
		}
	}
	return nil
}
