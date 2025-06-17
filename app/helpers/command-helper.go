package helpers

import (
	"os"
	"os/exec"
	"strings"
)

func GetPathEntry(method string) (string, bool) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", false
	}

	file, err := exec.LookPath(method)
	if err != nil {
		return "", false
	}

	return file, true
}

func SanetizeSingleQuotes(arguments string) (string, bool) {
	if strings.Count(arguments, "'")%2 == 0 {
		return strings.ReplaceAll(arguments, "'", ""), true
	}

	return arguments, false
}
