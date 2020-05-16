// Package task implements tasks for build stage
package task

import (
	"fmt"
	"io/ioutil"
)

// Task defines puppet for tasks
type Task interface {
	Do() error
	String() string
}

// MoveFile provides copy of the target file
// to the project
func MoveFile(inFile, outFile string) error {
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

// CreateFile provides creating of the file with data
func CreateFile(name string, data []byte) error {
	if name == "" {
		return fmt.Errorf("name of the file is not defined")
	}

	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to create file at: %s %v", outFile, err)
	}

	return nil
}
