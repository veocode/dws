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

func main() {
	cli := shell.NewCLIParser().Parse(os.Args)
	mapper := new(config.ActionMapper)

	actionName := cli.GetActionName()
	actionArgs := cli.GetActionArguments()
	actionData := repos.NewDataset()

	if !mapper.IsActionExists(actionName) {
		fmt.Printf("Unknown action: %s\n", actionName)
		os.Exit(1)
	}

	actionHandler := mapper.GetActionHandler(actionName)

	validateErr := actionHandler.Validate(actionArgs, actionData)
	if validateErr != nil {
		fmt.Printf("FAILED: %s\n", validateErr)
		os.Exit(ExitCodeBadCommand)
	}

	executeErr := actionHandler.Execute(actionArgs, actionData)
	if executeErr != nil {
		fmt.Printf("FAILED: %s\n", executeErr)
		os.Exit(ExitCodeError)
	}

	os.Exit(0)
}
