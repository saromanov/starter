package makefile

import (
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// Makefile implements create of the dockerfile
type Makefile struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Makefile{
		p: p,
	}
}

func (d *Makefile) String() string {
	return "create-makefile"
}

// Do defines action of the task
func (d *Makefile) Do() error {
	if d.p.Makefile == "" {
		return nil
	}

	if d.p.Makefile == "default" {
		return moveMakefile("./assets/makefile-default", "Makefile")
	}

	return moveMakefile(d.p.Makefile, "Makefile")
}

// moveMakefile provides moving of Makefile
func moveMakefile(inPath, outPath string) error {
	return task.MoveFile(inPath, outPath)
}
