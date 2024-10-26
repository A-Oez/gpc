package create

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"

	dbFactory "github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/structures/base"
	"github.com/pterm/pterm"
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
	projectNameFlag, err := cmd.Flags().GetString("p")
	databaseFlag, _ := cmd.Flags().GetString("db")
	openEditorFlag, _ := cmd.Flags().GetBool("code")

	if err != nil {
		fmt.Println("Error retrieving project name flag:", err)
		return
	}

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
	pterm.Success.Printf("Project %s successfully created!", projectNameFlag)
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
	
	dbType, isValid := dbFactory.ParseDatabaseType(databaseFlag)
	if !isValid {
		fmt.Printf("Given db input %s doesn't exist, no db structure created", databaseFlag)
		return
	}

	db := dbFactory.DatabaseServiceFactory(projectNameFlag, dbType)
	db.CreateDirectories()
	db.CreateFiles()
	db.UseCommand()
}