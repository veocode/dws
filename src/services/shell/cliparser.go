package shell

import "github.com/veocode/dws/src/repos"

type CLIParser struct {
	Arguments struct {
		Program string
		Action  string
		All     []string
	}
}

func NewCLIParser() *CLIParser {
	cp := new(CLIParser)
	return cp
}

func (cp *CLIParser) Parse(args []string) *CLIParser {
	cp.Arguments.Program = args[0]
	if len(args) == 1 {
		return cp
	}
	cp.Arguments.Action = args[1]
	cp.Arguments.All = args[1:]
	return cp
}

func (cp *CLIParser) GetActionName() string {
	return cp.Arguments.Action
}

func (cp *CLIParser) GetActionArgument() string {
	if len(cp.Arguments.All) == 0 {
		return ""
	}
	return cp.Arguments.All[0]
}

func (cp *CLIParser) GetActionArguments() *repos.Arguments {
	args := new(repos.Arguments)
	args.Set(cp.Arguments.All)
	return args
}
