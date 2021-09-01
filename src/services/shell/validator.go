package shell

import (
	"fmt"
	"os"
)

var ExecutableDeps = []string{
	"git",
	"docker",
	"docker-compose",
}

type Validator struct {
	targetDir string
}

func NewValidator(dir string) *Validator {
	if dir == "" {
		dir, _ = os.Getwd()
	}

	validator := new(Validator)
	validator.targetDir = dir
	return validator
}

func (v *Validator) CheckDeps() error {
	return CheckInstalledPrograms(ExecutableDeps...)
}

func (v *Validator) CheckDirForNewProject() error {
	if !IsPathExists(v.targetDir) {
		return fmt.Errorf("directory %s doesn't exist", v.targetDir)
	}

	isEmpty, err := IsDirEmpty(v.targetDir)
	if err != nil {
		return err
	}
	if !isEmpty {
		return fmt.Errorf("directory %s is not empty", v.targetDir)
	}

	return nil
}
