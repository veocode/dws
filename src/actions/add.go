package actions

import (
	"os"

	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/shell"
)

type Add struct {
}

func (action Add) Validate(args *repos.Arguments, data *repos.Dataset) error {
	targetDir, err := os.Getwd()
	if err != nil {
		return err
	}

	validator := shell.NewValidator(targetDir)
	err = validator.CheckDirContainsWorkspace()
	if err != nil {
		return err
	}

	return nil
}

func (action Add) Execute(args *repos.Arguments, data *repos.Dataset) error {
	return nil
}
