package actions

import (
	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/dws"
	"github.com/veocode/dws/src/services/shell"
)

type Update struct {
}

func (action Update) Validate(args *repos.Arguments, data *repos.Dataset) error {
	return shell.NewValidator("").CheckDeps()
}

func (action Update) Execute(args *repos.Arguments, data *repos.Dataset) error {
	shell.PrintOut("Rebuilding dsw hub cache...")

	hub, err := dws.NewHub()
	if err != nil {
		return err
	}

	err = hub.UpdateCache()
	if err != nil {
		return err
	}

	return nil
}
