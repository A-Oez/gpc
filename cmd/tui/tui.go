package tui

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "create standard Golang project - TUI",
	Run:   execute,
}

func init() {
	rootCmd.GetRootCmd().AddCommand(tuiCmd)
}

func execute(cmd *cobra.Command, args []string){	
	pterm.Info.Println("Starting TUI...")
}