package commands

import "os"

type cdCommand struct {
	path string
}

func CreateCdCommand(path string) *cdCommand {
	return &cdCommand{path: path}
}

func (cc *cdCommand) Execute() {
	err := os.Chdir(cc.path)
	if err != nil {
		panic("Error changing directory: " + err.Error())
	}
}
