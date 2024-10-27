package main

import (
	"github.com/A-Oez/gpc/cmd"
	_ "github.com/A-Oez/gpc/cmd/create"
	_ "github.com/A-Oez/gpc/cmd/tui"
)

func main() {
	cmd.Execute()
}
