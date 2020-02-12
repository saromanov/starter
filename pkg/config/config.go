package config

import (
	"fmt"

	"github.com/saromanov/cowrow"
	"github.com/saromanov/starter/pkg/models"
)

// Config provides definition of the project
type Config struct {
	Dockerfile string   `yaml:"dockerfile"`
	Makefile   string   `yaml:"makefile"`
	SubDirs    []string `yaml:"subdirs"`
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
	return &models.Project{
		Dockerfile: c.Dockerfile,
		Makefile:   c.Makefile,
		SubDirs:    c.SubDirs,
	}
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
