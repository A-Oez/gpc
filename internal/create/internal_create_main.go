package internal_create

import (
	"sync"
)

var (
	varDirName string
	varReadme  bool
	varCode    bool
)

func Execute(dirFlag string, rmdFlag bool, codeFlag bool) {
	initGlobalVariables(dirFlag, rmdFlag, codeFlag)
	var wg sync.WaitGroup

	mainDir()

	wg.Add(3)

	go func() {
		defer wg.Done()
		subDir()
	}()

	go func() {
		defer wg.Done()
		allfiles()
	}()

	go func() {
		defer wg.Done()
		allCommands()
	}()

	wg.Wait()
}

func initGlobalVariables(dirFlag string, rmdFlag bool, codeFlag bool) {
	varDirName = dirFlag
	varReadme = rmdFlag
	varCode = codeFlag
}
