package create

import (
	"errors"

	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/gpc/cmd"
	"github.com/A-Oez/gpc/internal"

	bp "github.com/A-Oez/gpc/internal/base"
	"github.com/pterm/pterm"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create golang projects",
	Run: execute,
}

func init() {
	var openEditor bool
	var api bool
	rootCmd.GetRootCmd().AddCommand(createCmd)
	createCmd.PersistentFlags().String("p", "", "define projectname")
	createCmd.PersistentFlags().String("db", "", "define database (optional)")
	createCmd.PersistentFlags().BoolVar(&openEditor, "code", false, "open project in code editor")
	createCmd.PersistentFlags().BoolVar(&api, "api", false, "create webservice")
}

func execute(cmd *cobra.Command, args []string){	
	projectNameFlag, err1 := cmd.Flags().GetString("p")
	databaseFlag, err2 := cmd.Flags().GetString("db")
	openEditorFlag, err3 := cmd.Flags().GetBool("code")
	apiFlag, err4 := cmd.Flags().GetBool("api")
	if err := errors.Join(err1, err2, err3, err4); err != nil{
		pterm.Error.Print("error retrieving flags:", err)
		return
	}

	if(apiFlag && databaseFlag != ""){
		pterm.Error.Print("invalid flags: you must choose only one — either '--api' or '--database'")
		return
	}

	bp := bp.BaseProject{
		ProjectName: projectNameFlag,
		OpenEditor: openEditorFlag,
	}
	internal.ExecuteCreation(bp, databaseFlag, apiFlag)
	pterm.Success.Printf("Project %s successfully created!", projectNameFlag)
}