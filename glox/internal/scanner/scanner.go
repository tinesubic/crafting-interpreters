package scanner

import (
	"crafting-interpreters/internal/util/loxerr"
	"fmt"
	"strconv"
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

func (s *Scanner) LogDebug(msg string) {
	loxerr.Debug(s.line, fmt.Sprintf("[Scanner] %s", msg))
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
		} else {
			s.LogDebug(fmt.Sprintf("Token: %v, Lexeme: %s, Literal %v", token.Type, token.Lexeme, token.Literal))
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
	switch c := s.advance(); c {
	case ' ', '\t', '\r':
		return nil // skip any whitespace
	case '(':
		return s.tokenize(LEFT_PAREN, nil)
	case ')':
		return s.tokenize(RIGHT_PAREN, nil)
	case '{':
		return s.tokenize(LEFT_BRACE, nil)
	case '}':
		return s.tokenize(RIGHT_BRACE, nil)
	case ',':
		return s.tokenize(COMMA, nil)
	case '.':
		return s.tokenize(DOT, nil)
	case '+':
		return s.tokenize(PLUS, nil)
	case '-':
		return s.tokenize(MINUS, nil)
	case '*':
		return s.tokenize(STAR, nil)
	case '\n':
		s.line++
		return nil
	case '!':
		if s.matchNext('=') {
			return s.tokenize(NOT_EQUAL, nil)
		} else {
			return s.tokenize(NOT, nil)
		}
	case '>':
		if s.matchNext('=') {
			return s.tokenize(GREATER_EQUAL, nil)
		} else {
			return s.tokenize(GREATER_THAN, nil)
		}
	case '<':
		if s.matchNext('=') {
			return s.tokenize(LESS_EQUAL, nil)
		} else {
			return s.tokenize(LESS_THAN, nil)
		}
	case '=':
		if s.matchNext('=') {
			return s.tokenize(EQUAL_EQUAL, nil)
		} else {
			return s.tokenize(EQUAL, nil)
		}
	case '/':
		if s.matchNext('/') {
			for s.peek() != NEWLINE && !s.isAtEnd() {
				_ = s.advance() // ignore comment data
			}
			return nil
		} else {
			return s.tokenize(SLASH, nil)
		}
	case '"':
		return s.parseString()
	default:
		if isDigit(c) {
			return s.parseNumber()
		} else if isAlpha(c) {
			return s.parseIdentifier()
		}
		loxerr.ReportError(s.line, fmt.Sprintf("Unexpected character %q at position %d", c, s.start))
	}
	return nil
}

func (s *Scanner) advance() byte {
	var next = s.source[s.current]
	s.current += 1
	return next
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	} else {
		return s.source[s.current]
	}
}

func (s *Scanner) tokenize(tokenType TokenType, value any) *Token {
	text := s.source[s.start:s.current]
	return &Token{
		Type:    tokenType,
		Line:    s.line,
		Literal: value,
		Lexeme:  text,
	}
}

func (s *Scanner) matchNext(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) parseString() *Token {
	for s.peek() != QUOTE && !s.isAtEnd() {
		if s.peek() == NEWLINE {
			s.line++
		}
		s.advance()
	}

	s.advance() // closing quote

	return &Token{
		Type:    STRING,
		Line:    s.line,
		Literal: s.source[s.start+1 : s.current-1],
		Lexeme:  s.source[s.start:s.current],
	}
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) parseNumber() *Token {
	for isDigit(s.peek()) || (s.peek() == DOT_RUNE && isDigit(s.peekNext())) {
		s.advance()
	}

	num, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		loxerr.ReportError(s.line, fmt.Sprintf("Could not parse %s as a number", s.source[s.start:s.current]))
	}

	return s.tokenize(NUMBER, num)
}

func (s *Scanner) parseIdentifier() *Token {
	for isAlphanumeric(s.peek()) {
		s.advance()
	}
	value := s.source[s.start:s.current]

	if tokenType, ok := reservedKeywords[value]; ok {
		return s.tokenize(tokenType, nil)
	} else {
		return s.tokenize(IDENTIFIER, value)
	}
}
