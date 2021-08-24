package repos

type Arguments struct {
	args []string
}

func (a *Arguments) Set(args []string) {
	a.args = args
}

func (a *Arguments) Exists() bool {
	return len(a.args) > 0
}

func (a *Arguments) Get(index int) string {
	if len(a.args)-1 < index {
		return ""
	}
	return a.args[index]
}

func (a *Arguments) GetOrDefault(index int, defaultValue string) string {
	value := a.Get(index)
	if value == "" {
		return defaultValue
	}
	return value
}
