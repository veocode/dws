package actions

import "fmt"

type Test struct {
}

func (t Test) Execute() {
	fmt.Println("Test")
}
