package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder

	var prevChar string
	var startsWithBackslash bool

	for _, char := range input {
		isBackslash := string(char) == "\\"
		isDigit := unicode.IsDigit(char)

		switch {
		case isBackslash && startsWithBackslash:
			prevChar = "\\"
			startsWithBackslash = false
		case isBackslash && prevChar == "":
			startsWithBackslash = true
		case isBackslash && prevChar != "":
			result.WriteString(prevChar)
			prevChar = ""
			startsWithBackslash = true

		case isDigit && startsWithBackslash:
			prevChar = string(char)
			startsWithBackslash = false
		case isDigit && prevChar == "":
			return "", ErrInvalidString
		case isDigit && prevChar != "":
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(prevChar, num))
			prevChar = ""

		case startsWithBackslash:
			return "", ErrInvalidString
		case prevChar == "":
			prevChar = string(char)
		case prevChar != "":
			result.WriteString(prevChar)
			prevChar = string(char)
		}
	}

	if startsWithBackslash {
		return "", ErrInvalidString
	}

	if prevChar != "" {
		result.WriteString(prevChar)
	}

	return result.String(), nil
}
