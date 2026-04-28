package aishit

import (
	"os"
	"os/exec"
)

func Execute(command string, args []string) {
	cmd := exec.Command(command, args...)

	// Direct the command's output/error streams to the current process's streams
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Optional: If the command needs input from your terminal
	cmd.Stdin = os.Stdin

	// Run() starts the command and waits for it to finish
	err := cmd.Run()
	if err != nil {
		// The error will likely be empty if you forwarded Stderr,
		// but you can still catch exit codes here.
		os.Exit(1)
	}
}
