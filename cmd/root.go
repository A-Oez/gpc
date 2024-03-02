package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-project",
	Short: "A program to spin up a quick Go project using a popular framework",
	Long: `Go Blueprint is a CLI tool that allows users to spin up a Go project with the corresponding structure seamlessly. 
It also gives the option to integrate with one of the more popular Go frameworks!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
