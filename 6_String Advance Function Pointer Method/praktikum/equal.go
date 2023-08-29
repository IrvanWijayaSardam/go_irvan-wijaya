package main

import (
	"fmt"
)

func Compare(a, b string) string {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	var commonSubstring string

	for i := 0; i < minLength; i++ {
		if a[i] == b[i] {
			commonSubstring += string(a[i])
		} else {
			break
		}
	}

	return commonSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
