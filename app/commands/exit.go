package commands

import (
	"os"
	"strconv"
)

type exitCommand struct {
	number int
}

func CreateExitCommand(query []string) *exitCommand {

	num, err := strconv.Atoi(query[0])
	if err != nil {
		panic("Invalid parameter for exit command, use 0 to exit")
	}

	return &exitCommand{number: num}
}

func (ec *exitCommand) Execute() {
	if ec.number != 0 {
		panic("Invalid parameter for exit command, use 0 to exit")
	}

	os.Exit(ec.number)
}
