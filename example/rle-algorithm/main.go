package main

import (
	"fmt"

	"github.com/sasakihasuto/compression-algorithms-lt/example/rle-algorithm/rle"
)

func main() {
	input := "aaabbbcccddeee"
	fmt.Println("Input:", input)
	encoded := rle.Encode(input)
	fmt.Println("Encoded:", encoded)

	decode, error := rle.Decode(encoded)
	if error != nil {
		fmt.Println("Error decoding:", error)
		return
	}
	fmt.Println("Decoded:", decode)
}
