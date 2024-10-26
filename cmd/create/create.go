package create

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"

	dbFactory "github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/structures/base"
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

	bp := bp.BaseProject{
		ProjectName: projectNameFlag, 
		OpenEditor: openEditorFlag,
	}
	bp.CreateMainDirectory()

	var wg sync.WaitGroup
	wg.Add(2)

	go createBaseStructure(bp, &wg)
	go createDBStructure(projectNameFlag, databaseFlag, &wg)

	wg.Wait()
}

func createBaseStructure(bp bp.BaseProject, wg *sync.WaitGroup){
	defer wg.Done()

	bp.CreateDirectories()
	bp.CreateFiles()
	bp.UseCommands()
}

func createDBStructure(projectNameFlag string, databaseFlag string, wg *sync.WaitGroup){
	defer wg.Done()

	if databaseFlag == ""{
		return
	}
	
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