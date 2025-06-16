package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type command interface {
	Execute()
}

type CommandType int

const (
	ExitCommand = iota
	EchoCommand
	TypeCommand
)

var commandType = map[CommandType]string{
	ExitCommand: "exit",
	EchoCommand: "echo",
	TypeCommand: "type",
}

var commandName = map[string]CommandType{
	"exit": ExitCommand,
	"echo": EchoCommand,
	"type": TypeCommand,
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
	default:
		method, found := getPathEntry(command)
		if !found {
			fmt.Printf("%s: command not found\n", command)
			return nil, errors.New("")
		}

		ec := CreateExternalCommnad(method, query)
		return ec, nil
	}
}

func getPathEntry(method string) (string, bool) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", false
	}

	file, err := exec.LookPath(method)
	if err != nil {
		return "", false
	}

	return file, true
}
