package commands

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/helpers"
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
		path, found := helpers.GetPathEntry(tc.method)
		if !found {
			fmt.Printf("%s: not found\n", tc.method)
			return
		}

		fmt.Printf("%v is %v\n", tc.method, path)
	}
}
