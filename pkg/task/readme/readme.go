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
	if d.p.EntryFile == "" {
		return nil
	}
	d1 := []byte(generate(d.p.EntryFile, d.p.Description, d.p.Author))
	err := ioutil.WriteFile(fmt.Sprintf("%s/README.md", d.p.Name), d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to write readme: %v", err)
	}
	return nil
}

func generate(name, description, author string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("# %s\n", name))
	if description != "" {
		builder.WriteString(fmt.Sprintf("%s\n", description))
	} else {
		builder.WriteString("\n")
	}
	builder.WriteString(fmt.Sprintf("### Author\n"))
	builder.WriteString(fmt.Sprintf("%s\n", author))
	builder.WriteString(fmt.Sprintf("### LICENCE\n"))
	return builder.String()
}
