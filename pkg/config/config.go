package config

import (
	"fmt"

	"github.com/saromanov/cowrow"
	"github.com/saromanov/starter/pkg/models"
)

// Config provides definition of the project
type Config struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	Dockerfile  string            `yaml:"dockerfile"`
	Makefile    string            `yaml:"makefile"`
	SubDirs     []string          `yaml:"subdirs"`
	Commands    []CommandLineSpec `yaml:"commands"`
	HostingPath string            `yaml:"hostingPath"`
	Badges      string            `yaml:"badges"`
	CI          []string          `yaml:"ci"`
}

// CommandLineSpec defines specification for command line
// For example it may contains command "add"
// Then it'll generate template for that command
type CommandLineSpec struct {
	Name string `yaml:"name"`
}

// Load provides loading of the config
func Load(path string) (*Config, error) {

	c := &Config{}
	if err := cowrow.LoadByPath(path, &c); err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}
	c.makeDefaults()
	return c, nil
}

// ToModel provides converting of the config to model
func (c *Config) ToModel() *models.Project {
	p := &models.Project{
		Name:       c.Name,
		Dockerfile: c.Dockerfile,
		Makefile:   c.Makefile,
		SubDirs:    c.SubDirs,
		Type:       models.StrToProjectType(c.Type),
	}
	if len(c.Commands) > 0 {
		for _, com := range c.Commands {
			p.Commands = append(p.Commands, models.Command{
				Name: com.Name,
			})
		}
	}
	return p
}

// makeDefaults provides filling of default config
func (c *Config) makeDefaults() {
	if len(c.SubDirs) == 0 {
		c.SubDirs = []string{"cmd", "pkg"}
	}
}

// DefaultConfig retruns default configuration for the project
func DefaultConfig() *Config {
	return &Config{
		SubDirs: []string{"cmd", "pkg"},
	}
}
