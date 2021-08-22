package actions

import (
	"fmt"
)

type Version struct {
}

func (v Version) Execute() {
	fmt.Println("Version: 0.0.1")
}
