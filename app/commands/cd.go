package commands

import (
	"fmt"
	"os"
)

type cdCommand struct {
	path string
}

func CreateCdCommand(path []string) *cdCommand {
	return &cdCommand{path: path[0]}
}

func (cc *cdCommand) Execute() {
	if cc.path == "~" {
		path, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("Error getting user home directory: %v", err))
		}

		os.Chdir(path)
	} else {
		err := os.Chdir(cc.path)
		if err != nil {
			fmt.Printf("cd: %v: No such file or directory\n", cc.path)
		}
	}
}
