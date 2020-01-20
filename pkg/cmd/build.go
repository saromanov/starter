package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/saromanov/starter/pkg/models"
)

// Build provides building of the tree structure for project
func Build() error {

}

func consoleRead() (*models.Project, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add definition of the project")
	fmt.Println("---------------------")
	fmt.Println("Name of the project")
	p := &models.Project{}
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read user input: %v", err)
	}
	p.Name = name

	fmt.Println("Author of the project")
	author, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read user input: %v", err)
	}
	p.Name = author
	p.Author = author

	return p, nil
}

// validateDirectories provides validation of the dirs
// from input with project-layout
func validateDirectories(dirs []string) ([]string, error) {
	return dirs, nil
}
