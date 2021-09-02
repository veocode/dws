package actions

import (
	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/shell"
)

type Version struct {
}

func (action Version) Validate(args *repos.Arguments, data *repos.Dataset) error {
	return nil
}

func (action Version) Execute(args *repos.Arguments, data *repos.Dataset) error {
	shell.PrintOut("Version: 0.0.0")
	return nil
}
