package models

type Workspace struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
	Dir    string
}

func NewWorkspace(targetDir string) *Workspace {
	workspace := new(Workspace)
	workspace.Dir = targetDir
	return workspace
}
