package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var escapeableCharsInDoubleQuotes = []rune{'"', '\\', '$', '`'}
var quotes = []rune{'"', '\''}

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

func SanetizeCommand(command string) (method string, arguments []string) {
	command = strings.TrimSpace(command)
	sanetizedCommand := sanetize(command)

	// fmt.Printf("sanetizedMethod: %v\n", strings.Join(sanetizedCommand, ", "))
	// fmt.Printf("method: %v\n", sanetizedCommand[0])
	// fmt.Printf("arguments[0:]: %v\n", strings.Join(sanetizedCommand[1:], ", "))
	return sanetizedCommand[0], sanetizedCommand[1:]
}

func SanetizeMethod(method string) string {
	fmt.Printf("method: %v\n", method)
	sanetizedMethod := sanetize(method)
	return sanetizedMethod[0]
}

func SanetizeArguments(arguments string) []string {
	arguments = strings.TrimSpace(arguments)
	return sanetize(arguments)
}

func sanetize(arguments string) []string {
	var tokens strings.Builder
	var values []string
	isInSingleQuote := false
	isInDoubleQuote := false
	isBackslash := false

	push := func() {
		if tokens.Len() > 0 {
			values = append(values, tokens.String())
			tokens.Reset()
		}
	}

	handleBackslash := func(r rune) bool {
		if !isBackslash {
			return false
		}

		if (isInDoubleQuote && !slices.Contains(escapeableCharsInDoubleQuotes, r)) || isInSingleQuote {
			tokens.WriteRune('\\')
		}

		tokens.WriteRune(r)
		isBackslash = false
		return true
	}

	handleDoubleQuote := func(r rune) bool {
		if !isInDoubleQuote {
			return false
		}

		if r == '\'' {
			tokens.WriteRune(r)
			return true
		}

		return false
	}

	handleSingleQuote := func(r rune) bool {
		if !isInSingleQuote {
			return false
		}

		if r == '"' {
			tokens.WriteRune(r)
			return true
		}

		return false
	}

	for _, char := range arguments {
		if handled := handleBackslash(char); handled {
			continue
		}

		if handled := handleDoubleQuote(char); handled {
			continue
		}

		if handled := handleSingleQuote(char); handled {
			continue
		}

		switch char {
		case '\\':
			isBackslash = true
		case '\'':
			if !isInDoubleQuote {
				isInSingleQuote = !isInSingleQuote
			}
		case '"':
			if !isInSingleQuote {
				isInDoubleQuote = !isInDoubleQuote
			}
		default:
			if char == ' ' && (isInSingleQuote || isInDoubleQuote) { // char is whitespace and is inside quotes so we append it
				tokens.WriteRune(' ')
			} else if char == ' ' { // char is whitespace but not inside quotes -> new argument
				push()
			} else {
				tokens.WriteRune(char) // char is not whitespace
			}
		}
	}
	push()

	return values
}
