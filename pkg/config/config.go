package config

import (
	"fmt"
	"time"

	"github.com/saromanov/cowrow"
)

// Config provides definition of the project
type Config struct {
	Name       string   `yaml:"name"`
	Author     string   `yaml:"author"`
	GitPath    string   `yaml:"gitpath"`
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

	if c.Name == "" {
		return nil, fmt.Errorf("name is not defined")
	}

	c.makeDefaults()
	return c, nil
}

// makeDefaults provides filling of default config
func (c *Config) makeDefaults() {
	if len(c.Subdirs) == 0 {
		c.Subdirs = []string{"cmd", "pkg"}
	}
}

