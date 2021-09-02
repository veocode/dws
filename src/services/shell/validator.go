package shell

import (
	"fmt"
	"path/filepath"

	"github.com/veocode/dws/src/config"
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
	validator := new(Validator)
	validator.targetDir = dir
	return validator
}

func (v *Validator) CheckDeps() error {
	return CheckInstalledPrograms(ExecutableDeps...)
}

func (v *Validator) CheckDirForNewWorkspace() error {
	err := v.CheckDeps()
	if err != nil {
		return err
	}

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

func (v *Validator) CheckDirContainsWorkspace() error {
	dwsFilePath := filepath.Join(v.targetDir, config.FileMetaWorkspace)
	if !IsPathExists(dwsFilePath) {
		return fmt.Errorf("directory %s doesn't contain a dws workspace", v.targetDir)
	}
	return nil
}
