package main

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

func handleCommand(command string) {
	// remove trailing newline character
	command = command[:len(command)-1]

	parameters := strings.Fields(command)

	if len(parameters) == 0 {
		fmt.Println("No command entered")
		return
	}

	method := parameters[0]

	query, found := strings.CutPrefix(command, method)
	if !found {
		fmt.Printf("%s: command not found\n", method)
		return
	}

	if commandHandle, _ := commands.CreateCommand(method, query); commandHandle != nil {
		commandHandle.Execute()
	}
}
