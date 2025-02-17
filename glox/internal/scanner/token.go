package scanner

import "fmt"

type TokenType int

const (
	NEWLINE = '\n'
)

const (
	EOF       TokenType = iota
	SEMICOLON TokenType = iota

	// math ops
	MINUS TokenType = iota
	PLUS  TokenType = iota
	STAR  TokenType = iota
	SLASH TokenType = iota

	// Logical ops
	GREATER_THAN  TokenType = iota
	LESS_THAN     TokenType = iota
	EQUAL         TokenType = iota
	LESS_EQUAL    TokenType = iota
	GREATER_EQUAL TokenType = iota
	EQUAL_EQUAL   TokenType = iota
	NOT_EQUAL     TokenType = iota
	NOT

	// Chars
	LEFT_PAREN  TokenType = iota
	RIGHT_PAREN TokenType = iota
	LEFT_BRACE  TokenType = iota
	RIGHT_BRACE TokenType = iota
	COMMA       TokenType = iota
	DOT         TokenType = iota

	IDENTIFIER TokenType = iota
	STRING     TokenType = iota
	NUMBER     TokenType = iota

	AND    TokenType = iota
	OR     TokenType = iota
	IF     TokenType = iota
	ELSE   TokenType = iota
	FALSE  TokenType = iota
	TRUE   TokenType = iota
	FOR    TokenType = iota
	WHILE  TokenType = iota
	NIL    TokenType = iota
	CLASS  TokenType = iota
	FUN    TokenType = iota
	RETURN TokenType = iota
	SUPER  TokenType = iota
	THIS   TokenType = iota
	VAR    TokenType = iota
)

type Token struct {
	Type    TokenType
	Line    int
	Literal any
	Lexeme  string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s:%v", t.Lexeme, t.Literal)
}
