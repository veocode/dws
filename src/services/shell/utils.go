package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/otiai10/copy"
)

func IsProgramInstalled(program string) bool {
	command := exec.Command(program, "--help")
	err := command.Run()
	return err == nil
}

func CheckInstalledPrograms(programs ...string) error {
	for _, program := range programs {
		if !IsProgramInstalled(program) {
			return fmt.Errorf("%s not found, please install", program)
		}
	}
	return nil
}

func IsDirEmpty(dirPath string) (bool, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Execute(appWithArgs ...string) error {
	cmd := exec.Command(appWithArgs[0], appWithArgs[1:]...)
	err := cmd.Run()
	return err
}

func CopyDirectory(src, dst string) error {
	return copy.Copy(src, dst)
}
