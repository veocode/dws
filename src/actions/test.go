package actions

import (
	"fmt"

	"github.com/veocode/dws/src/repos"
)

type Test struct {
}

func (action Test) Validate(args *repos.Arguments) error {
	return nil
}

func (action Test) Execute(args *repos.Arguments) error {
	fmt.Println("Test")
	return nil
}
