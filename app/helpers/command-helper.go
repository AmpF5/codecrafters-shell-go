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

func SanetizeArguments(arguments string) []string {
	arguments = strings.TrimSpace(arguments)
	return sanetize(arguments)
}

func sanetize(arguments string) []string {
	var tokens strings.Builder
	var values []string

	push := func() {
		if tokens.Len() > 0 {
			values = append(values, tokens.String())
			tokens.Reset()
		}
	}

	isInSingleQuote := false
	isInDoubleQuote := false
	isBackslash := false

	for _, char := range arguments {
		if char == '\\' && !isInSingleQuote && !isInDoubleQuote {
			isBackslash = !isBackslash
			continue
		}

		if char == '\'' && !isInDoubleQuote {
			isInSingleQuote = !isInSingleQuote
			continue
		}

		if char == '"' && !isInSingleQuote && !isBackslash {
			isInDoubleQuote = !isInDoubleQuote
			continue
		}

		if isBackslash {
			tokens.WriteRune(char)
			isBackslash = false
			continue
		}

		if char == ' ' && (isInSingleQuote || isInDoubleQuote) {
			tokens.WriteRune(' ')
			continue
		} else if char == ' ' {
			push()
		} else {
			tokens.WriteRune(char)
		}
	}
	push()

	return values
}
