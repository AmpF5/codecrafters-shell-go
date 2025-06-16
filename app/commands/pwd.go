package commands

import (
	"fmt"
	"os"
)

type pwdCommand struct{}

func CreatePwdCommand() *pwdCommand {
	return &pwdCommand{}
}

func (pd *pwdCommand) Execute() {
	cDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Error getting current directory: %v", err))
	}

	fmt.Printf("%v\n", cDir)
}
