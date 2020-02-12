package entryfile

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// Entryfile defines for the entry file task
type Entryfile struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Entryfile{
		p: p,
	}
}

// String returns representation of the task
func (d *Entryfile) String() string {
	return "entry-file"
}

// Do defines action of the task
func (d *Entryfile) Do() error {
	if d.p.EntryFile == "" || d.p.EntryFile == "\n" {
		return nil
	}
	d1 := []byte(fmt.Sprintf("package %s", d.p.Name))
	fileName := d.getFileName()
	err := ioutil.WriteFile(fileName, d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to write file %s: %v", fileName, err)
	}
	return nil
}

// getFileName returns result file name
func (d *Entryfile) getFileName() string {
	if strings.HasSuffix(d.p.EntryFile, ".go") {
		return fmt.Sprintf("%s/%s", d.p.Name, d.p.EntryFile)
	}
	return fmt.Sprintf("%s/%s.go", d.p.Name, d.p.EntryFile)
}
