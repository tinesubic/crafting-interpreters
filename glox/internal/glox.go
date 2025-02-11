package internal

import (
	"crafting-interpreters/internal/loxerr"
	"crafting-interpreters/internal/scanner"
	"fmt"
)

type Glox struct {
	hadError bool
}

func (g *Glox) Run(input string) error {
	scanner := scanner.NewScanner(input)
	tokens, err := scanner.ScanTokens()
	if err != nil {
		return err
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	if g.hadError {
		return loxerr.GloxError{}
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
