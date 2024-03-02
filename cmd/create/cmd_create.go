/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"github.com/spf13/cobra"

	"projectcreator/cli/cmd"

	create "projectcreator/cli/internal/create"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		dirFlag, _ := cmd.Flags().GetString("dir")
		rmdFlag, _ := cmd.Flags().GetBool("rmd")
		codeFlag, _ := cmd.Flags().GetBool("code")

		create.Execute(dirFlag, rmdFlag, codeFlag)
	},
}

func init() {
	cmd.GetRootCmd().AddCommand(createCmd)
	cmd.GetRootCmd().PersistentFlags().String("dir", "", "flag to define projectname/directory")

	var rmd bool
	cmd.GetRootCmd().PersistentFlags().BoolVar(&rmd, "rmd", false, "")

	var code bool
	cmd.GetRootCmd().PersistentFlags().BoolVar(&code, "code", false, "")
}
