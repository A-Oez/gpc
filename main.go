package main

import (
	"github.com/A-Oez/GoProjectCreator/cmd"
	_ "github.com/A-Oez/GoProjectCreator/cmd/create"
	_ "github.com/A-Oez/GoProjectCreator/cmd/tui"
)

func main() {
	cmd.Execute()
}
