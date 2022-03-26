package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var unpackedStr strings.Builder
	unpackedStr.Reset()
	for i, _rune := range s {
		if s == "" {
			return "", nil
		}
		if unicode.IsDigit(_rune) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(_rune) && unicode.IsDigit(rune(s[i-1])) {
			return "", ErrInvalidString
		}
		if i < len(s)-1 {
			if unicode.IsLetter(_rune) && !unicode.IsDigit(rune(s[i+1])) {
				unpackedStr.WriteRune(_rune)
			}
		}
		if i == len(s)-1 {
			if unicode.IsLetter(_rune) && !unicode.IsDigit(rune(s[len(s)-1])) {
				unpackedStr.WriteRune(_rune)
			}
		}
		if unicode.IsDigit(_rune) {
			_s := string(s[i-1])
			_count, _ := strconv.Atoi(string(_rune))
			unpackedStr.WriteString(strings.Repeat(_s, _count))
		}
	}
	return unpackedStr.String(), nil
}
