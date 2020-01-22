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
	
	return c, nil
}

