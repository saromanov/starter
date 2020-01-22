package project

import (
	"fmt"
	"io/ioutil"
)

// moveFile provides copy of the target file
// to the project
func moveFile(inFile, outFile string) error {
	input, err := ioutil.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("unable to open file: %s %v", inFile, err)
	}

	err = ioutil.WriteFile(outFile, input, 0644)
	if err != nil {
		return fmt.Errorf("unable to create file at: %s %v", outFile, err)
	}

	return nil
}
