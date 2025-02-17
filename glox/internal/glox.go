package internal

import (
	"crafting-interpreters/internal/scanner"
	"crafting-interpreters/internal/util/loxerr"
)

type Glox struct {
	hadError bool
}

func (g *Glox) Run(input string) error {
	scanner := scanner.NewScanner(input)
	_, err := scanner.ScanTokens()
	if err != nil {
		return err
	}

	if g.hadError {
		return loxerr.GloxError{
			Phase: loxerr.ScannerPhase,
		}
	}

	return nil
}

func (g *Glox) Error(line int, msg string) {
	loxerr.ReportError(line, msg)
	g.hadError = true
}

func (g *Glox) ResetError() {
	g.hadError = false
}
