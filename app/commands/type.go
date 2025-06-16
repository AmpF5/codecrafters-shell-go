package commands

import (
	"fmt"
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
		path, found := getPathEntry(tc.method)
		if !found {
			panic(fmt.Sprintf("%s: command not found", tc.method))
		}

		fmt.Printf("%v is %v\n", tc.method, path)
	}
}
