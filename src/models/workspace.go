package models

func NewWorkspace() *Workspace {
	return new(Workspace)
}

type Workspace struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
}
