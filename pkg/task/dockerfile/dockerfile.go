package dockerfile

import (
	"io/ioutil"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

// Dockerfile implements create of the dockerfile
type Dockerfile struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Dockerfile{
		p: p,
	}
}

func (d *Dockerfile) String() string {
	return "create-dockerfile"
}

// Do defines action of the task
func (d *Dockerfile) Do() error {
	if d.Dockerfile == "default" {
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
