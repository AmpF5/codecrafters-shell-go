package commands

import (
	"fmt"
	"os"
	"strings"
)

type typeCommand struct {
	method string
}

func CreateTypeCommand(query string) *typeCommand {
	params := strings.Fields(query)

	if len(params) != 1 {
		panic("Invalid parameter for type command, use a single word")
	}

	return &typeCommand{method: params[0]}
}

func (tc *typeCommand) Execute() {
	if _, exists := commandName[tc.method]; exists {
		fmt.Printf("%v is a shell builtin\n", tc.method)
	} else {
		file, notFound := getPathEntries(tc.method)
		if notFound {
			fmt.Printf("%s: not found\n", tc.method)
		} else {
			fmt.Printf("%v is %v", tc.method, file)
		}
	}
}

func getPathEntries(method string) (string, bool) {
	path := os.Getenv("PATH")
	// path, ok := os.LookupEnv("PATH")
	if path == "" {
		panic("PATH environment variable is not set or empty")
	}

	for dir := range strings.SplitSeq(path, string(os.PathListSeparator)) {
		file := dir + "/" + method
		if _, err := os.Stat(file); err == nil {
			return file, false
		}
	}

	return "", true
}
