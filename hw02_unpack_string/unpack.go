package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here
	var prevRune rune
	var err error
	var isEscaped int
	res := []rune{}
	for i, el := range s {
		if unicode.IsDigit(el) {
			res, err = workDigit(prevRune, &isEscaped, res, el, i)
			if err != nil {
				return "", err
			}
		} else {
			res, err = workOthers(el, prevRune, &isEscaped, res, i)
			if err != nil {
				return "", err
			}
		}
		if isEscaped != 0 && isEscaped < i {
			isEscaped = 0
		}
		prevRune = el
	}
	return string(res), nil
}

func workOthers(el int32, prevRune rune, isEscaped *int, res []rune, i int) ([]rune, error) {
	switch el {
	case '\\':
		if prevRune == '\\' && *isEscaped == 0 {
			res = append(res, el)
			*isEscaped = i
		}
		return res, nil
	default:
		if prevRune == '\\' {
			return res, ErrInvalidString
		}
		res = append(res, el)
		return res, nil
	}
}

func workDigit(prevRune rune, isEscaped *int, res []rune, el int32, i int) ([]rune, error) {
	switch {
	case prevRune == 0:
		return res, ErrInvalidString
	case prevRune == '\\' && *isEscaped == 0:
		res = append(res, el)
		*isEscaped = i
		return res, nil
	case unicode.IsDigit(prevRune):
		if *isEscaped != 0 {
			res, err := repeatRune(el, res, prevRune)
			if err != nil {
				return res, fmt.Errorf("%c is not digit", el)
			}
			return res, nil
		}
		return res, ErrInvalidString

	default:
		res, err := repeatRune(el, res, prevRune)
		if err != nil {
			return res, fmt.Errorf("%c is not digit", el)
		}
		return res, nil
	}
}

func repeatRune(el int32, res []rune, prevRune rune) ([]rune, error) {
	cnt, err := strconv.Atoi(string(el))
	if cnt == 0 {
		res = res[:len(res)-1]
	}
	for i := 1; i < cnt; i++ {
		res = append(res, prevRune)
	}
	return res, err
}
