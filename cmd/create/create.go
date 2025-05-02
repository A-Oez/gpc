package create

import (
	"errors"
	"log"

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
	var webservice bool
	rootCmd.GetRootCmd().AddCommand(createCmd)
	createCmd.PersistentFlags().String("p", "", "define projectname")
	createCmd.PersistentFlags().String("db", "", "define database (optional)")
	createCmd.PersistentFlags().BoolVar(&openEditor, "code", false, "open project in code editor")
	createCmd.PersistentFlags().BoolVar(&webservice, "webservice", false, "create webservice")
}

func execute(cmd *cobra.Command, args []string){	
	projectNameFlag, err1 := cmd.Flags().GetString("p")
	databaseFlag, err2 := cmd.Flags().GetString("db")
	openEditorFlag, err3 := cmd.Flags().GetBool("code")
	webserviceFlag, err4 := cmd.Flags().GetBool("webservice")
	if err := errors.Join(err1, err2, err3, err4); err != nil{
		log.Fatal("error retrieving flags:", err)
	}

	bp := bp.BaseProject{
		ProjectName: projectNameFlag,
		OpenEditor: openEditorFlag,
	}
	internal.ExecuteCreation(bp, databaseFlag, webserviceFlag)
	pterm.Success.Printf("Project %s successfully created!", projectNameFlag)
}