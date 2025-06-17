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

func SanetizeArguments(arguments string) string {
	arguments = strings.TrimSpace(arguments)
	arguments = sanetizeMultipleSpaces(arguments)
	arguments, _ = sanetizeSingleQuotes(arguments)

	return arguments
}

func sanetizeSingleQuotes(arguments string) (string, bool) {
	if strings.Count(arguments, "'")%2 == 0 {
		return strings.ReplaceAll(arguments, "'", ""), true
	}

	return arguments, false
}

func sanetizeMultipleSpaces(arguments string) string {
	return strings.Join(strings.Fields(arguments), " ")
}
