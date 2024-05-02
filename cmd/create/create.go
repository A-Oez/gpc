/*
Copyright © 2024 A.Oez
*/
package create

import (
	"github.com/spf13/cobra"

	"github.com/A-Oez/GoProjectCreator/cmd"

	bp "github.com/A-Oez/GoProjectCreator/internal/base_project"
	_ "github.com/A-Oez/GoProjectCreator/internal/db"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create standard Golang project",
	Run: execute,
}

func init() {
	var openEditor bool
	cmd.GetRootCmd().AddCommand(createCmd)
	createCmd.PersistentFlags().String("p", "", "define projectname")
	createCmd.PersistentFlags().String("db", "", "define database (optional)")
	createCmd.PersistentFlags().BoolVar(&openEditor, "code", false, "open project in code editor")
}

func execute(cmd *cobra.Command, args []string){
	projectNameFlag, _ := cmd.Flags().GetString("p")
	//databaseFlag, _ := cmd.Flags().GetBool("db")
	openEditorFlag, _ := cmd.Flags().GetBool("code")

	//create base project
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: openEditorFlag,
	}
	bp.CreateMainDirectory()
	
	//ab hier kann alles in go routinen laufen -> vielleicht in channel schreiben wenn main directory erstellt wurde? 
	bp.CreateDirectories()
	bp.CreateFiles()
	bp.UseCommands()

/*
	if databaseFlag{
		db := db.DatabaseProject{
			ProjectName: projectNameFlag, 
		}

		db.CreateDirectories()
	}	
*/
}