package commands

import (
	"fmt"
	"os"
)

type cdCommand struct {
	path string
}

func CreateCdCommand(path string) *cdCommand {
	return &cdCommand{path: path}
}

func (cc *cdCommand) Execute() {
	err := os.Chdir(cc.path)
	if err != nil {
		fmt.Printf("cd: %v: No such file or directory\n", cc.path)
	}
}
