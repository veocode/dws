package contracts

import "github.com/veocode/dws/src/repos"

type Action interface {
	Validate(args *repos.Arguments, data *repos.Dataset) error
	Execute(args *repos.Arguments, data *repos.Dataset) error
}
