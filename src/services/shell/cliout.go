package shell

import (
	"fmt"
	"os"
)

func PrintOut(pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, pattern+"\n", args)
}

func PrintErr(pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern+"\n", args)
}

func PrintDone() {
	PrintOut("SUCCESS: Done")
}
