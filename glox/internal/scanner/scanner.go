package scanner

type Scanner struct {
	source string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	return nil, nil
}
