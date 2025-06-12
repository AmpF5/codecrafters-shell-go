package commands

import "fmt"

type echoCommand struct {
	value string
}

func CreateEchoCommand(query string) *echoCommand {
	return &echoCommand{value: query}
}

func (ec echoCommand) Execute() {
	fmt.Println(ec.value)
}
