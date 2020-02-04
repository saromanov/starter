package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/saromanov/starter/pkg/config"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/project"
)

// Build provides building of the tree structure for project
func Build(projectFlag, configPath string) error {
	pr := &models.Project{}
	if configPath != "" {
		cfg, err := config.Load(configPath)
		if err != nil {
			return errors.Wrap(err, "unable to load config")
		}
		pr = cfg.ToModel()
	}
	err := consoleRead(pr)
	if err != nil {
		return errors.Wrap(err, "unable to read data from console")
	}
	pr.Type = models.StrToProjectType(projectFlag)
	if err := project.Build(pr); err != nil {
		return fmt.Errorf("unable to build project: %v", err)
	}

	return nil
}

func consoleRead(p *models.Project) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add definition of the project")
	fmt.Println("---------------------")
	name, err := readLine(reader, "Name of the project")
	if err != nil {
		return fmt.Errorf("unable to read user input: %v", err)
	}
	if len(name) < 3 || len(name) > 30 {
		return fmt.Errorf("unable to validate name of the project length")
	}
	if name == "" {
		return errNoName
	}
	p.Name = name[:len(name)-1]

	author, err := readLine(reader, "Author of the project")
	if err != nil {
		return fmt.Errorf("unable to read user input: %v", err)
	}
	if author == "" {
		return errNoAuthor
	}
	p.Author = author

	gitPath, err := readLine(reader, "Git path(optional)")
	if err != nil {
		return fmt.Errorf("unable to read git path: %v", err)
	}
	p.GitPath = gitPath

	entryFile, err := readLine(reader, "Entry file(optional)")
	if err != nil {
		return fmt.Errorf("unable to read git path: %v", err)
	}

	p.EntryFile = entryFile

	return nil
}

func readLine(reader *bufio.Reader, name string) (string, error) {
	fmt.Println(name)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to read line: %v", err)
	}
	return line, nil
}

// validateDirectories provides validation of the dirs
// from input with project-layout
func validateDirectories(dirs []string) ([]string, error) {
	return dirs, nil
}
