package actions

import (
	"fmt"
	"os"

	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/shell"
)

var executableDeps = []string{
	"git",
	"docker",
	"docker-compose",
}

type Init struct {
}

func (action Init) Validate(args *repos.Arguments) error {
	return shell.CheckInstalledPrograms(executableDeps...)
}

func (action Init) Execute(args *repos.Arguments) error {

	// folderPath = shell.CLIParser.Get

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("init in " + cwd)

	return nil
}
