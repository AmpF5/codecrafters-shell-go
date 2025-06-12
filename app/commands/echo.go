package commands

import (
	"fmt"
	"strings"
)

type echoCommand struct {
	value string
}

func CreateEchoCommand(query string) *echoCommand {
	query = strings.TrimLeft(query, " ")
	return &echoCommand{value: query}
}

func (ec echoCommand) Execute() {
	fmt.Println(ec.value)
}
