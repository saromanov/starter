package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const envVariable = "STARTER_TEMPLATES"

// ListTemplates returns list of templates for starter
// it should get from directory defined at STARTER_TEMPLATES
// environment variable
func ListTemplates() error {
	dir := os.Getenv(envVariable)
	if dir == "" {
		return fmt.Errorf("%s is not defined", envVariable)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("unable to get files from directory: %s", dir)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fmt.Println(path.Ext(f.Name()))
	}
	return nil
}
