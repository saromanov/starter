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
	d1 := []byte(generate(d.p))
	err := ioutil.WriteFile(fmt.Sprintf("%s/README.md", d.p.Name), d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to write readme: %v", err)
	}
	return nil
}

// generate provides generating of the README
func generate(p *models.Project) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("# %s\n", p.Name))
	plugins := []Plugin{description}
	if len(p.Badges) > 0 {
		plugins = append(plugins, addBadges)
	}
	for _, plugin := range plugins {
		plugin(&builder, p)
	}
	builder.WriteString(fmt.Sprintf("### Author\n"))
	builder.WriteString(fmt.Sprintf("%s\n", p.Author))
	builder.WriteString(fmt.Sprintf("### LICENCE\n"))
	return builder.String()
}
