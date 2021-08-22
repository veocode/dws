package main

import (
	"fmt"
	"os"

	"github.com/veocode/dws/src/config"
	"github.com/veocode/dws/src/services"
)

func main() {

	cli := services.NewCLIParser().Parse(os.Args)
	actionName := cli.GetActionName()
	actionMapper := new(config.ActionMapper)

	if !actionMapper.IsActionExists(actionName) {
		fmt.Printf("Unknown action: %s\n", actionName)
		os.Exit(1)
	}

	actionHandler := actionMapper.GetActionHandler(actionName)
	actionHandler.Execute()

	os.Exit(0)

}
