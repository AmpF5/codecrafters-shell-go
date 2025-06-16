package commands

import (
	"os"
	"os/exec"
	"strings"
)

type externalCommand struct {
	method    string
	arguments string
}

func CreateExternalCommnad(method, arguments string) *externalCommand {
	methodFromPath := strings.Split(method, "/")
	return &externalCommand{
		method:    methodFromPath[len(methodFromPath)-1],
		arguments: arguments,
	}
}

func (ec *externalCommand) Execute() {
	cmd := exec.Command(ec.method, strings.Fields(ec.arguments)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	// if err != nil {
	// 	panic("Cannot execute external command: " + ec.method + " " + ec.arguments)
	// }

	// fmt.Printf("%s\n", out)
}
