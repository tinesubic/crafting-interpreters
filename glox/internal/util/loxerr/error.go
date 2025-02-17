package loxerr

import (
	"fmt"
	"os"
)

type Phase string

const ScannerPhase Phase = "scanner"
const RuntimePhase Phase = "runtime"

type GloxError struct {
	Phase Phase
	Line  int
}

func (g GloxError) Error() string {
	return fmt.Sprintf("[L%d] GLOX has encountered a problem in %s phase", g.Line, g.Phase)
}

func ReportError(line int, message string) {
	report(line, "Error", "", message)
}

func Debug(line int, message string) {
	if _, err := fmt.Fprintf(os.Stdout, "[DEBUG] [%d] %s\n", line, message); err != nil {
		panic(err)
	}
}

func report(line int, what string, where string, msg string) {
	if _, err := fmt.Fprintf(os.Stderr, "[ERROR] [%d] %s %s: %s\n", line, what, where, msg); err != nil {
		panic(err)
	}
}
