package scanner

import (
	"crafting-interpreters/internal/util/loxerr"
	"fmt"
)

type Scanner struct {
	source string
	tokens []Token

	start   int // starting position of lexeme
	current int // current position in source
	line    int // current line
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	for {
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		token := s.scanToken()
		if token == nil {
			continue
		}
		s.tokens = append(s.tokens, *token)
	}

	eofToken := Token{Type: EOF}

	s.tokens = append(s.tokens, eofToken)

	return s.tokens, nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() *Token {
	c := s.advance()
	switch c {
	case " ", "\t", "\r":
		return nil
	case "(":
		return s.tokenize(LEFT_PAREN)
	case ")":
		return s.tokenize(RIGHT_PAREN)
	case "{":
		return s.tokenize(LEFT_BRACE)
	case "}":
		return s.tokenize(RIGHT_BRACE)
	case ",":
		return s.tokenize(COMMA)
	case ".":
		return s.tokenize(DOT)
	case "+":
		return s.tokenize(PLUS)
	case "-":
		return s.tokenize(MINUS)
	case "/":
		return s.tokenize(SLASH)
	case "*":
		return s.tokenize(STAR)
	case "\n":
		s.line++
		return nil
	default:
		loxerr.ReportError(s.line, fmt.Sprintf("Unexpected character %q at position %d\n", c, s.start))
	}
	return nil
}

func (s *Scanner) advance() string {
	var next = s.source[s.current]
	s.current += 1
	return string(next)
}

func (s *Scanner) tokenize(tokenType TokenType) *Token {
	text := s.source[s.start:s.current]
	return &Token{
		Type:    tokenType,
		Line:    s.line,
		Literal: nil,
		Lexeme:  text,
	}
}
