package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/project"
)

// Build provides building of the tree structure for project
func Build(projectFlag string) error {
	pr, err := consoleRead()
	if err != nil {
		return err
	}
	pr.Type = models.StrToProjectType(projectFlag)
	if err := project.Build(pr); err != nil {
		return fmt.Errorf("unable to build project: %v", err)
	}

	return nil
}

func consoleRead() (*models.Project, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add definition of the project")
	fmt.Println("---------------------")
	p := &models.Project{}
	fmt.Println("Name of the project")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read user input: %v", err)
	}
	if len(name) < 3 || len(name) > 30 {
		return nil, fmt.Errorf("unable to validate name of the project length")
	}
	if name == "" {
		return nil, errNoName
	}
	p.Name = name[:len(name)-1]

	fmt.Println("Author of the project")
	author, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read user input: %v", err)
	}
	if author == "" {
		return nil, errNoAuthor
	}
	p.Author = author

	fmt.Println("Git path")
	gitPath, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read git path: %v", err)
	}
	p.GitPath = gitPath

	return p, nil
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
