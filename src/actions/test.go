package actions

import (
	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/shell"
)

type Test struct {
}

func (action Test) Validate(args *repos.Arguments, data *repos.Dataset) error {
	return nil
}

func (action Test) Execute(args *repos.Arguments, data *repos.Dataset) error {
	shell.PrintOut("TEST")
	return nil
}
