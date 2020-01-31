package gitremote

import (
	"github.com/pkg/errors"
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
	if err := exec.Run(d.p.Name, "git", "init"); err != nil {
		return errors.Wrap(err, "unable to run 'git init' command")
	}
	if err := exec.Run(d.p.Name, "git", "remote", "add", "origin", d.p.GitPath); err != nil {
		return errors.Wrap(err, "unable to run 'git remote add origin'")
	}
	return nil
}
