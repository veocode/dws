package shell

import (
	"fmt"
	"os"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func PrintOut(pattern string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprintf(os.Stdout, pattern+"\n", args...)
	} else {
		fmt.Fprintln(os.Stdout, pattern)
	}
}

func PrintErr(pattern string, args ...interface{}) {
	pattern = ColorRed + "FAILED: " + ColorReset + pattern
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, pattern+"\n", args...)
	} else {
		fmt.Fprintln(os.Stderr, pattern)
	}
}

func PrintDone() {
	pattern := ColorGreen + "SUCCESS: " + ColorReset + "Done"
	PrintOut(pattern)
}
