package scanner

import "unicode"

func isDigit(d byte) bool {
	r := rune(d)
	return r >= '0' && r <= '9'
}

func isAlpha(d byte) bool {
	r := rune(d)
	return unicode.IsLetter(r) || r == '_'
}

func isAlphanumeric(d byte) bool {
	return isAlpha(d) || isDigit(d)
}
