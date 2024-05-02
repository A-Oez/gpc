package create

import (
	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"

	bp "github.com/A-Oez/GoProjectCreator/internal/base_project"
	db "github.com/A-Oez/GoProjectCreator/internal/db"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create standard Golang project",
	Run: execute,
}

func init() {
	var openEditor bool
	rootCmd.GetRootCmd().AddCommand(createCmd)
	createCmd.PersistentFlags().String("p", "", "define projectname")
	createCmd.PersistentFlags().String("db", "", "define database (optional)")
	createCmd.PersistentFlags().BoolVar(&openEditor, "code", false, "open project in code editor")
}

func execute(cmd *cobra.Command, args []string){
	projectNameFlag, _ := cmd.Flags().GetString("p")
	databaseFlag, _ := cmd.Flags().GetString("db")
	openEditorFlag, _ := cmd.Flags().GetBool("code")

	//create base project
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: openEditorFlag,
	}
	bp.CreateMainDirectory()
	//TODO: go routines ?
	bp.CreateDirectories()
	bp.CreateFiles()
	bp.UseCommands()


	if databaseFlag != ""{ //TODO: add method to check if given db type exists -> db enum?
		db := db.DatabaseProject{
			ProjectName: projectNameFlag, 
		}
		db.CreateDirectories()
		db.CreateFiles()
	}	

}