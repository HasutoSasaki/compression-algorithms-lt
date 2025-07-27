package rle

import "strconv"

func Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var encoded string
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			encoded += string(input[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}
	encoded += string(input[len(input)-1]) + strconv.Itoa(count)
	return encoded
}
