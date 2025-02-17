package scanner

import "testing"

func TestAdvance(t *testing.T) {
	s := Scanner{
		source:  "12",
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    0,
	}

	c := s.advance()
	if c != '1' {
		t.Error("Expected byte 1")
	}
	c = s.advance()
	if c != '2' {
		t.Error("Expected byte 2")
	}
}

func TestScanToken(t *testing.T) {
	s := Scanner{
		source:  "-\n+//123",
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    0,
	}

	token := s.scanToken()
	if token.Type != MINUS {
		t.Error()
	}
	token = s.scanToken()
	if token != nil {
		t.Error()
	}
	token = s.scanToken()
	if token.Type != PLUS {
		t.Error()
	}
	token = s.scanToken()
	if token != nil {
		t.Error()
	}
}
