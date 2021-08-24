package actions

import (
	"fmt"

	"github.com/veocode/dws/src/repos"
)

type Version struct {
}

func (action Version) Validate(args *repos.Arguments) error {
	return nil
}

func (action Version) Execute(args *repos.Arguments) error {
	fmt.Println("Version: 0.0.1")
	return nil
}
