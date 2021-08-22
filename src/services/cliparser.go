package services

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
