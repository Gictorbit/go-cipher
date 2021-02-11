package sezar

import (
	"errors"
	"unicode"
)

type Caesar struct {
	Text  string
	Shift int32
}

func (sezar Caesar) EncodeText() string {

	res := ""
	for _, letter := range sezar.Text {
		if unicode.IsUpper(letter) {
			tmpchar := modulus((toIndex(letter) + int(sezar.Shift)), 26)
			mainchar, _ := toChar(tmpchar, false)
			res += string(mainchar)
		} else if unicode.IsLower(letter) {
			tmpchar := modulus((toIndex(letter) + int(sezar.Shift)), 26)
			mainchar, _ := toChar(tmpchar, true)
			res += string(mainchar)
		} else {
			res += string(letter)
		}

	}
	return res
}

func (sezar Caesar) DecodeText() string {
	res := ""
	for _, letter := range sezar.Text {
		if unicode.IsUpper(letter) {
			tmpchar := modulus((toIndex(letter) - int(sezar.Shift)), 26)
			mainchar, _ := toChar(tmpchar, false)
			res += string(mainchar)
		} else if unicode.IsLower(letter) {
			tmpchar := modulus((toIndex(letter) - int(sezar.Shift)), 26)
			mainchar, _ := toChar(tmpchar, true)
			res += string(mainchar)
		} else {
			res += string(letter)
		}

	}
	return res
}

func toChar(lettNom int, isLower bool) (rune, error) {
	if lettNom > 25 || lettNom < 0 {
		return '-', errors.New("charachter out of range english letters 0-25")
	}
	if isLower {
		return rune('a' + lettNom), nil
	} else {
		return rune('A' + lettNom), nil
	}
}

func toIndex(letter rune) int {
	if unicode.IsLower(letter) {
		return int(letter - rune('a'))
	} else {
		return int(letter - rune('A'))
	}
}

func modulus(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
