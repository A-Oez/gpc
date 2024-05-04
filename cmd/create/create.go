package create

import (
	"fmt"

	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"

	bp "github.com/A-Oez/GoProjectCreator/internal/structures/base"
	dbFactory "github.com/A-Oez/GoProjectCreator/internal/structures/db"
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

	createBaseStructure(projectNameFlag, openEditorFlag, databaseFlag)
	createDBStructure(projectNameFlag, databaseFlag)
}

func createBaseStructure(projectNameFlag string, openEditorFlag bool, databaseFlag string){
	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: openEditorFlag,
	}
	bp.CreateMainDirectory()
	bp.CreateDirectories()
	bp.CreateFiles()
	bp.UseCommands()
}

func createDBStructure(projectNameFlag string, databaseFlag string){
	if dbType, isValid := dbFactory.ParseDatabaseType(databaseFlag); isValid{  
		db := dbFactory.DatabaseServiceFactory(projectNameFlag, dbType)
		db.CreateDirectories()
		db.CreateFiles()
		db.UseCommand()
	} else {
		message := fmt.Sprintf("Given db input %s doesnt exist, no db structure created", databaseFlag)
		fmt.Println(message)
	}	
}