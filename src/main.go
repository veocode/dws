package main

import (
	"os"

	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services"
	"github.com/veocode/dws/src/services/shell"
)

const ExitCodeBadCommand = 65
const ExitCodeError = 1
const ExitCodeSuccess = 0

func main() {
	cliParser := shell.NewCLIParser().Parse(os.Args)
	router := services.NewRouter()

	actionName := cliParser.GetActionName()
	actionArgs := cliParser.GetActionArguments()

	if !router.IsActionExists(actionName) {
		shell.PrintErr("Unknown action: %s", actionName)
		os.Exit(ExitCodeError)
	}

	actionData := repos.NewDataset()
	actionHandler := router.GetActionHandler(actionName)

	validateErr := actionHandler.Validate(actionArgs, actionData)
	if validateErr != nil {
		shell.PrintErr("%s", validateErr)
		os.Exit(ExitCodeBadCommand)
	}

	executeErr := actionHandler.Execute(actionArgs, actionData)
	if executeErr != nil {
		shell.PrintErr("%s", executeErr)
		os.Exit(ExitCodeError)
	}

	shell.PrintDone()
	os.Exit(ExitCodeSuccess)
}
