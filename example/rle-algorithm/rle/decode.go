package rle

import (
	"strconv"
	"strings"
)

func Decode(encoded string) (string, error) {
	var decoded strings.Builder

	for i := 0; i < len(encoded); i += 2 {
		if i+1 >= len(encoded) {
			break
		}

		char := encoded[i]
		count, err := strconv.Atoi(string(encoded[i+1]))
		if err != nil {
			return "", err
		}

		decoded.WriteString(strings.Repeat(string(char), count))
	}

	return decoded.String(), nil
}
