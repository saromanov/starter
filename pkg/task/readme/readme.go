package readme

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/task"
)

var errNoName = errors.New("name of the project is not defined")

// Readme defines for the readme task
type Readme struct {
	p *models.Project
}

// New creates of the task
func New(p *models.Project) task.Task {
	return &Readme{
		p: p,
	}
}

// String returns representation of the task
func (d *Readme) String() string {
	return "readme"
}

// Do defines action of the task
func (d *Readme) Do() error {
	d1 := []byte(generate(d.p.Name, d.p.Description))
	err := ioutil.WriteFile(fmt.Sprintf("%d/README.md", d.p.Name), d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to write readme: %v", err)
	}
	return nil
}

func generate(name, description, author string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("# %s", name))
	if description != "" {
		builder.WriteString(description)
	}
	builder.WriteString(fmt.Sprintf("### Author"))
	builder.WriteString(author)
	return builder.String()
}
