package gitremote

import (
	"fmt"

	"github.com/saromanov/starter/pkg/exec"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// Gitremote defines for the entry file task
type Gitremote struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Gitremote{
		p: p,
	}
}

// String returns representation of the task
func (d *Gitremote) String() string {
	return "git-remote"
}

// Do defines action of the task
func (d *Gitremote) Do() error {
	if d.p.GitPath == "" {
		return nil
	}
	if err := exec.Run(d.p.Name, "git", "remote", "add", "origin", d.p.GitPath); err != nil {
		return fmt.Errorf("unable to run 'go mod init' command: %v", err)
	}
	return nil
}
