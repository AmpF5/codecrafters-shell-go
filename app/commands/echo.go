package commands

import (
	"fmt"
	"strings"
)

type echoCommand struct {
	value string
}

func CreateEchoCommand(query []string) *echoCommand {
	return &echoCommand{value: strings.Join(query, " ")}
}

func (ec echoCommand) Execute() {
	fmt.Println(ec.value)
}
