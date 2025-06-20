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

	method, arguments := helpers.SanetizeCommand(command)

	// method := helpers.SanetizeMethod(command)

	// arguments, found := strings.CutPrefix(command, method)
	// if !found {
	// 	fmt.Printf("%s: command not found\n", method)
	// 	return
	// }

	// argumentsSanetized := helpers.SanetizeArguments(arguments)

	if commandHandle, _ := commands.CreateCommand(method, arguments); commandHandle != nil {
		commandHandle.Execute()
	}
}
