package loxerr

import (
	"fmt"
	"os"
)

type GloxError struct{}

func (g GloxError) Error() string {
	return "GLOX has encountered a problem"
}

func ReportError(line int, message string) {
	report(line, "Error", "", message)
}

func report(line int, what string, where string, msg string) {
	if _, err := fmt.Fprintf(os.Stderr, "[%s] %s %s: %s", line, what, where, msg); err != nil {
		panic(err)
	}
}
