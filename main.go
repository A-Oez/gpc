package main

import (
	"projectcreator/cli/cmd"

	_ "projectcreator/cli/cmd/create"
)

func main() {
	cmd.Execute()
}
