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

func (bc *typeCommand) Execute() {
	if _, exists := commandName[bc.method]; exists {
		fmt.Printf("%v is a shell builtin\n", bc.method)
	} else {
		fmt.Printf("%s: not found\n", bc.method)
	}
}
