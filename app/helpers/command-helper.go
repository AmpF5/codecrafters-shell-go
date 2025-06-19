package helpers

import (
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
