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
	for i, r := range s {
		if s == "" {
			return "", nil
		}
		if unicode.IsDigit(r) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(r) && unicode.IsDigit(rune(s[i-1])) {
			return "", ErrInvalidString
		}
		if i < len(s)-1 {
			if unicode.IsLetter(r) && !unicode.IsDigit(rune(s[i+1])) {
				unpackedStr.WriteRune(r)
			}
		}
		if i == len(s)-1 {
			if unicode.IsLetter(r) && !unicode.IsDigit(rune(s[len(s)-1])) {
				unpackedStr.WriteRune(r)
			}
		}
		if unicode.IsDigit(r) {
			s := string(s[i-1])
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			unpackedStr.WriteString(strings.Repeat(s, count))
		}
	}
	return unpackedStr.String(), nil
}
