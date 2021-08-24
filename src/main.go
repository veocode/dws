package main

import (
	"fmt"
	"os"

	"github.com/veocode/dws/src/config"
	"github.com/veocode/dws/src/services/shell"
)

func main() {
	cli := shell.NewCLIParser().Parse(os.Args)
	actionName := cli.GetActionName()
	actionMapper := new(config.ActionMapper)
	actionArgs := cli.GetActionArguments()

	if !actionMapper.IsActionExists(actionName) {
		fmt.Printf("Unknown action: %s\n", actionName)
		os.Exit(1)
	}

	actionHandler := actionMapper.GetActionHandler(actionName)

	validateErr := actionHandler.Validate(actionArgs)
	if validateErr != nil {
		fmt.Printf("FAILED: %s", validateErr)
		os.Exit(65)
	}

	executeErr := actionHandler.Execute(actionArgs)
	if executeErr != nil {
		fmt.Printf("FAILED: %s", executeErr)
		os.Exit(1)
	}

	os.Exit(0)
}
