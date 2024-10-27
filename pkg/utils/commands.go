package utils

import (
	"bytes"
	"log"
	"os/exec"
)

func ExecuteCommand(cmd *exec.Cmd) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute command: %v\nStderr: %s", err, stderr.String())
	}
}