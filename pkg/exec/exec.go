// Package exec implements running of console commands
package exec

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Run provides running of the command
func Run(command string, flags ...string) error {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}
	return nil
}
