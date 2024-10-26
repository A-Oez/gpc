package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteCommand(cmd *exec.Cmd) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("exception:", err)
		fmt.Println("exception message:", stderr.String())
		return
	}
}
