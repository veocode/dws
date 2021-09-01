package actions

import (
	"fmt"
	"path/filepath"

	"github.com/veocode/dws/src/repos"
	"github.com/veocode/dws/src/services/hub"
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
	err := shell.CheckInstalledPrograms(executableDeps...)
	if err != nil {
		return err
	}

	targetDir, err := action.getTargetDir(args)
	if err != nil {
		return err
	}

	if !shell.IsPathExists(targetDir) {
		return fmt.Errorf("directory %s doesn't exist", targetDir)
	}

	isEmpty, err := shell.IsDirEmpty(targetDir)
	if err != nil {
		return err
	}
	if !isEmpty {
		return fmt.Errorf("directory %s is not empty", targetDir)
	}

	data.Set("targetDir", targetDir)

	return nil
}

func (action Init) Execute(args *repos.Arguments, data *repos.Dataset) error {
	targetDir := data.GetString("targetDir")
	fmt.Printf("Initializing new workspace in %s...\n", targetDir)

	hub, err := hub.NewHub()
	if err != nil {
		return err
	}

	err = hub.ExtractPath("init/scratch", targetDir)
	if err != nil {
		return err
	}

	return nil
}

func (action Init) getTargetDir(args *repos.Arguments) (string, error) {
	return filepath.Abs(args.GetOrDefault(0, "."))
}
