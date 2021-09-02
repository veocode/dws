package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Ask(prompt string, defaultValue interface{}) (response string, err error) {
	fmt.Fprint(os.Stdout, formatPrompt(prompt, defaultValue))
	response, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	response = strings.TrimSpace(strings.Replace(response, "\n", "", -1))
	if response == "" && defaultValue != nil {
		response = fmt.Sprintf("%v", defaultValue)
	}

	return
}

func formatPrompt(prompt string, defaultValue interface{}) string {
	pattern := ""
	pattern += fmt.Sprint(ColorBlue)
	pattern += "  >>"
	pattern += fmt.Sprint(ColorGreen)
	pattern += " %s"
	pattern += fmt.Sprint(ColorReset)
	if defaultValue == nil || defaultValue == "" {
		pattern += ": "
		return fmt.Sprintf(pattern, prompt)
	}
	pattern += " ["
	pattern += fmt.Sprint(ColorYellow)
	pattern += "%v"
	pattern += fmt.Sprint(ColorReset)
	pattern += "]: "
	return fmt.Sprintf(pattern, prompt, defaultValue)
}
