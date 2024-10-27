package tui

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"

	rootCmd "github.com/A-Oez/GoProjectCreator/cmd"
	"github.com/A-Oez/GoProjectCreator/internal"
	bp "github.com/A-Oez/GoProjectCreator/internal/config/base"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "create golang projects - TUI",
	Run:   execute,
}

var (
	databaseOptions = []string{"Postgres"} 
)

func init() {
	rootCmd.GetRootCmd().AddCommand(tuiCmd)
}

func execute(cmd *cobra.Command, args []string){
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	defer area.Stop()

	setupHeader()
	projectName := setProjectName()
	dbStr := setupDBCreation(projectName)

	//choose if editor should be opened
	pterm.Info.Printfln("Should the project be opened in an editor?")
	editor,_ := pterm.DefaultInteractiveConfirm.Show()

	bp := bp.BaseProject{
		ProjectName: projectName,
		OpenEditor: editor,
	}

	internal.ExecuteCreation(bp, dbStr)
	pterm.Println()
	pterm.Success.Printf("Project %s successfully created!", projectName)
}

func setupHeader(){
	//Header
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("GP", pterm.FgLightMagenta.ToStyle()),
		putils.LettersFromStringWithStyle("C", pterm.FgCyan.ToStyle()),
	).Render()
	pterm.DefaultHeader.Println("GoProjectCreator ~ A TUI for fast and flexible Go project creation ~ made with pterm")
	pterm.Println()
}

func setProjectName() string {
	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("Set project name").Show()
	pterm.Println()
	return result
}

func setupDBCreation(projectName string) string{
	pterm.Println()
	pterm.Info.Printfln("Would you like to set up a database for the project: %s?", projectName)
	dbYes, _ := pterm.DefaultInteractiveConfirm.Show()
	pterm.Println()

	var db string
	if dbYes {
		db = selectDatabase()
		pterm.Info.Printfln("Selected database: %s", pterm.Green(db))
	}

	pterm.Println()
	return db
}

func selectDatabase() string {
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(databaseOptions).Show()
	if err != nil {
		log.Fatalf("Error selecting database: %v", err)
	}
	return selectedOption
}