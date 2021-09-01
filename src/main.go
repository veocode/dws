package main

import (
	"fmt"
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
		fmt.Fprintf(os.Stderr, "Unknown action: %s\n", actionName)
		os.Exit(ExitCodeError)
	}

	actionData := repos.NewDataset()
	actionHandler := router.GetActionHandler(actionName)

	validateErr := actionHandler.Validate(actionArgs, actionData)
	if validateErr != nil {
		fmt.Fprintf(os.Stderr, "FAILED: %s\n", validateErr)
		os.Exit(ExitCodeBadCommand)
	}

	executeErr := actionHandler.Execute(actionArgs, actionData)
	if executeErr != nil {
		fmt.Fprintf(os.Stderr, "FAILED: %s\n", executeErr)
		os.Exit(ExitCodeError)
	}

	fmt.Fprintf(os.Stdout, "SUCCESS: done\n")
	os.Exit(ExitCodeSuccess)
}
