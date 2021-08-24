package shell

import (
	"fmt"
	"os/exec"
)

func IsProgramInstalled(program string) bool {
	command := exec.Command(program, "--help")
	err := command.Run()
	return err == nil
}

func CheckInstalledPrograms(programs ...string) error {
	for _, program := range programs {
		if !IsProgramInstalled(program) {
			return fmt.Errorf("%s not found", program)
		}
	}
	return nil
}
