package internal_create

import (
	"bytes"
	"fmt"
	"os/exec"
)

func allCommands() {
	cmdMod := exec.Command("go", "mod", "init", varDirName)
	cmdMod.Dir = varDirName
	executeCommand(cmdMod)

	if varCode {
		cmdCode := exec.Command("code", ".")
		cmdCode.Dir = varDirName
		executeCommand(cmdCode)
	}
}

func executeCommand(cmd *exec.Cmd) {
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
