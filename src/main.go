package main

import (
	"os"

	"github.com/veocode/dws/src/config"
	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/shell"
)

const ExitCodeBadCommand = 65
const ExitCodeError = 1
const ExitCodeSuccess = 0

func main() {
	cli := shell.NewCLIParser().Parse(os.Args)
	router := new(config.ActionRouter)

	actionName := cli.GetActionName()
	actionArgs := cli.GetActionArguments()

	if !router.IsActionExists(actionName) {
		shell.PrintErr("Unknown action: %s", actionName)
		os.Exit(ExitCodeError)
	}

	actionData := repos.NewDataset()
	actionHandler := router.GetActionHandler(actionName)

	validateErr := actionHandler.Validate(actionArgs, actionData)
	if validateErr != nil {
		shell.PrintErr("FAILED: %s", validateErr)
		os.Exit(ExitCodeBadCommand)
	}

	executeErr := actionHandler.Execute(actionArgs, actionData)
	if executeErr != nil {
		shell.PrintErr("FAILED: %s", executeErr)
		os.Exit(ExitCodeError)
	}

	shell.PrintDone()
	os.Exit(ExitCodeSuccess)
}
