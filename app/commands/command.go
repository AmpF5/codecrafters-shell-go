package commands

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/helpers"
)

type command interface {
	Execute()
}

type CommandType int

const (
	ExitCommand = iota
	EchoCommand
	TypeCommand
	PwdCommand
	CdCommand
)

var commandType = map[CommandType]string{
	ExitCommand: "exit",
	EchoCommand: "echo",
	TypeCommand: "type",
	PwdCommand:  "pwd",
	CdCommand:   "cd",
}

var commandName = map[string]CommandType{
	"exit": ExitCommand,
	"echo": EchoCommand,
	"type": TypeCommand,
	"pwd":  PwdCommand,
	"cd":   CdCommand,
}

func CreateCommand(command, query string) (command, error) {
	switch command {
	case "exit":
		ec := CreateExitCommand(query)
		return ec, nil
	case "echo":
		ec := CreateEchoCommand(query)
		return ec, nil
	case "type":
		bc := CreateTypeCommand(query)
		return bc, nil
	case "pwd":
		pwdc := CreatePwdCommand()
		return pwdc, nil
	case "cd":
		cc := CreateCdCommand(query)
		return cc, nil
	default:
		method, found := helpers.GetPathEntry(command)
		if !found {
			fmt.Printf("%s: command not found\n", command)
			return nil, errors.New("")
		}

		ec := CreateExternalCommnad(method, query)
		return ec, nil
	}
}
