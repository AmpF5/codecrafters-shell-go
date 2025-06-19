package helpers

import (
	"os"
	"os/exec"
	"slices"
	"strings"
)

var escapeableCharsInDoubleQuotes = []rune{'"', '\\', '$', '`'}

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

	handleBackslash := func(r rune) {
		if !isBackslash {
			return
		}

		if isInDoubleQuote && !slices.Contains(escapeableCharsInDoubleQuotes, r) {
			tokens.WriteRune('\\')
		}

		isBackslash = false
	}

	for _, char := range arguments {
		handleBackslash(char)

		switch char {
		case '\\':
			isBackslash = true
		case '\'':
			if isInDoubleQuote {
				tokens.WriteRune(char)
			} else {
				isInSingleQuote = !isInSingleQuote
			}
		case '"':
			isInDoubleQuote = !isInDoubleQuote
		default:
			if char == ' ' && (isInSingleQuote || isInDoubleQuote) { // char is whitespace and is inside quotes so we append it
				tokens.WriteRune(' ')
			} else if char == ' ' { // char is whitespace but not inside quotes -> new argument
				push()
			} else {
				tokens.WriteRune(char) // char is not whitespace
			}
		}
		// if char == '\\' && !isInSingleQuote && !isBackslash {
		// 	isBackslash = !isBackslash

		// 	if isInDoubleQuote {
		// 		continue
		// 	}
		// }

		// if char == '\'' && !isInDoubleQuote && !isBackslash {
		// 	isInSingleQuote = !isInSingleQuote
		// 	continue
		// }

		// if char == '"' && !isInSingleQuote && !isBackslash {
		// 	isInDoubleQuote = !isInDoubleQuote
		// 	continue
		// }

		// if isBackslash {
		// 	tokens.WriteRune(char)
		// 	isBackslash = false
		// 	continue
		// }

		// if char == ' ' && (isInSingleQuote || isInDoubleQuote) {
		// 	tokens.WriteRune(' ')
		// 	continue
		// } else if char == ' ' {
		// 	push()
		// } else {
		// 	tokens.WriteRune(char)
		// }
	}
	push()

	return values
}
