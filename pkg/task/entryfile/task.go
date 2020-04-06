package entryfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	subPath, err := d.createSubDirs()
	if err != nil {
		return fmt.Errorf("unable to create subdirs: %v", err)
	}
	fileName := d.getFileName(subPath)
	d1 := []byte(fmt.Sprintf("package %s", d.getPackageName(fileName)))
	if err = ioutil.WriteFile(fileName, d1, 0644); err != nil {
		return fmt.Errorf("unable to write file %s: %v", fileName, err)
	}
	return nil
}

// createSubDirs provides checking of the file
// if input data contains subdirs, then create it
// its returns result path
func (d *Entryfile) createSubDirs() (string, error) {
	subDirs := path.Dir(d.p.EntryFile)
	if subDirs == "." {
		return d.p.EntryFile, nil
	}
	if strings.HasPrefix(subDirs, "./") {
		subDirs = subDirs[2:len(subDirs)]
	}
	return fmt.Sprintf("%s%s", d.p.Name, d.p.EntryFile), os.MkdirAll(fmt.Sprintf("%s%s", d.p.Name, subDirs), 0777)
}

// getFileName returns result file name
func (d *Entryfile) getFileName(resultPath string) string {
	if strings.HasSuffix(resultPath, ".go") {
		return resultPath
	}
	return fmt.Sprintf("%s.go", resultPath)
}

// getPackageName returns package name
func (d *Entryfile) getPackageName(resultPath string) string {
	dirs, _ := path.Split(resultPath)
	result := strings.Split(dirs, "/")
	if len(result) == 1 {
		return result[1]
	}
	return result[len(result)-2]
}
