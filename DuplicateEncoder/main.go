package main

import (
	"fmt"
	"strings"
)

func DuplicateEncoder(word string) (encoded string) {
	for _, char := range word {
		if strings.Count(strings.ToLower(word), strings.ToLower(string(char))) > 1 {
			encoded += ")"
		} else {
			encoded += "("
		}
	}

	return encoded
}

func main() {
	fmt.Printf("%s\n", DuplicateEncoder("din"))
	fmt.Printf("%s\n", DuplicateEncoder("recede"))
	fmt.Printf("%s\n", DuplicateEncoder("Success"))
}
