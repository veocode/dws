package actions

import (
	"fmt"

	"github.com/veocode/dws/src/repos"
)

type Add struct {
}

func (action Add) Validate(args *repos.Arguments, data *repos.Dataset) error {
	currentDir

	fmt.Printf("Current dir: ")
	return nil
}

func (action Add) Execute(args *repos.Arguments, data *repos.Dataset) error {
	return nil
}
