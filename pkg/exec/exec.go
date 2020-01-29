// Package exec implements running of console commands
package exec

import (
	"fmt"
	"os/exec"
)

// Run provides running of the command
func Run(dir, command string, flags ...string) error {
	cmd := exec.Command(command, flags...)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}
	return nil
}
