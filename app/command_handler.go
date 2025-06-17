package main

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/helpers"
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

	arguments, found := strings.CutPrefix(command, method)
	if !found {
		fmt.Printf("%s: command not found\n", method)
		return
	}

	arguments = strings.TrimSpace(arguments)
	san, _ := helpers.SanetizeSingleQuotes(arguments)

	arguments = san

	if commandHandle, _ := commands.CreateCommand(method, arguments); commandHandle != nil {
		commandHandle.Execute()
	}
}
