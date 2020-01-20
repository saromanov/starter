package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// Build provides building of the tree structure for project
func Build() error {

}

func consoleRead() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add definition of the project")
	fmt.Println("---------------------")
	fmt.Println("Name of the project")
	text, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read user input: %v", err)
	}
}
