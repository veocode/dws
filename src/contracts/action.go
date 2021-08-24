package contracts

import "github.com/veocode/dws/src/repos"

type Action interface {
	Validate(args *repos.Arguments) error
	Execute(args *repos.Arguments) error
}
