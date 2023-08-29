package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	result := ""

	for _, char := range input {
		if char == ' ' {
			result += " "
		} else {
			charCode := int(char)
			shiftedCode := (charCode-97+offset)%26 + 97
			result += string(shiftedCode)
		}
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cncv
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
