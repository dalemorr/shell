package main

import (
	"errors"
	"strings"
)

func decodeOctal(s string) ([]byte, error) {
	if !isOctal(s) {
		return *new([]byte), errors.New("invalid octal string")
	}

	str := strings.ToLower(s)
	data := make([]byte, len(s))

	for i, char := range str {
		if char >= 0x30 && char <= 0x39 {
			data[i] = byte(char - 0x30)
		} else if char >= 0x61 && char <= 0x66 {
			data[i] = byte(char - 0x57)
		}
	}

	return data, nil
}

func isOctal(s string) bool {
	str := strings.ToLower(s)

	for _, char := range str {
		if !((char >= 0x30 && char <= 0x39) || (char >= 0x61 && char <= 0x66)) {
			return false
		}
	}

	return true
}

func octalToHex(b1, b2 byte) byte {
	return 0x10*b1 + b2
}
