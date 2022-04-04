package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	rs := []rune(s)
	var us strings.Builder
	us.Reset()
	if len(rs) == 0 {
		return s, nil
	}
	if unicode.IsDigit(rs[0]) {
		return "", ErrInvalidString
	}
	for i, r := range rs {
		if unicode.IsDigit(r) && unicode.IsDigit(rs[i-1]) {
			return "", ErrInvalidString
		}
		if i < len(rs)-1 {
			if !unicode.IsDigit(r) && !unicode.IsDigit(rs[i+1]) {
				us.WriteRune(r)
			}
		}
		if i == len(rs)-1 {
			if !unicode.IsDigit(r) && !unicode.IsDigit(rs[len(rs)-1]) {
				us.WriteRune(r)
			}
		}
		if unicode.IsDigit(r) {
			s := rs[i-1]
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			us.WriteString(strings.Repeat(string(s), count))
		}
	}
	return us.String(), nil
}
