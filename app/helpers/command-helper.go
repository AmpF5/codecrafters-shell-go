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

func sanetizeSingleQuotes(arguments string) (string, bool) {
	if strings.Count(arguments, "'")%2 == 0 {
		return strings.ReplaceAll(arguments, "'", ""), true
	}

	return arguments, false
}

func sanetizeMultipleSpaces(arguments string) string {
	// singleQuotesValues := strings.FieldsFunc(arguments, isSingleQuote)
	// strings.FieldsFuncSeq
	return strings.Join(strings.FieldsFunc(arguments, isSingleQuote), "")
}

func isSingleQuote(c rune) bool {
	return c == '\''
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
	for _, char := range arguments {
		if char == '\'' {
			isInSingleQuote = !isInSingleQuote
			continue
		}

		if char == ' ' && isInSingleQuote {
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
