package commands

import (
	"os"
	"os/exec"
	"strings"
)

type externalCommand struct {
	method    string
	arguments []string
}

func CreateExternalCommnad(method string, arguments []string) *externalCommand {
	methodFromPath := strings.Split(method, "/")
	return &externalCommand{
		method:    methodFromPath[len(methodFromPath)-1],
		arguments: arguments,
	}
}

func (ec *externalCommand) Execute() {
	cmd := exec.Command(ec.method, ec.arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
