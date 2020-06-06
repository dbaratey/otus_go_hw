package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
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
			switch {
			case prevRune == 0:
				return "", ErrInvalidString
			case prevRune == '\\' && isEscaped == 0:
				res = append(res, el)
				isEscaped = i
			case unicode.IsDigit(prevRune):
				if isEscaped != 0 {
					res, err = repeatRune(el, res, prevRune)
					if err != nil {
						log.Fatal(string(el), " is not digit")
					}
				} else {
					return "", ErrInvalidString
				}

			default:
				res, err = repeatRune(el, res, prevRune)
				if err != nil {
					log.Fatal(string(el), " is not digit")
				}
			}
		} else {
			switch el {
			case '\\':
				if prevRune == '\\' && isEscaped == 0 {
					res = append(res, el)
					isEscaped = i
				}
			default:
				if prevRune == '\\' {
					return "", ErrInvalidString
				}
				res = append(res, el)
			}
		}
		if isEscaped != 0 && isEscaped < i {
			isEscaped = 0
		}
		prevRune = el
	}
	return string(res), nil
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
