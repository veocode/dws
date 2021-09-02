package actions

import (
	"path/filepath"

	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/dws"
	"github.com/veocode/dws/src/services/shell"
)

var executableDeps = []string{
	"git",
	"docker",
	"docker-compose",
}

type Init struct {
}

func (action Init) Validate(args *repos.Arguments, data *repos.Dataset) error {
	targetDir, err := action.getTargetDir(args)
	if err != nil {
		return err
	}

	err = shell.NewValidator(targetDir).CheckDirForNewWorkspace()
	if err != nil {
		return err
	}

	data.Set("targetDir", targetDir)
	return nil
}

func (action Init) Execute(args *repos.Arguments, data *repos.Dataset) error {
	targetDir := data.GetString("targetDir")
	shell.PrintOut("Initializing new workspace in %s...", targetDir)
	return dws.NewWorkspaceManager().CreateWorkspace(targetDir)
}

func (action Init) getTargetDir(args *repos.Arguments) (string, error) {
	return filepath.Abs(args.GetOrDefault(0, "."))
}
