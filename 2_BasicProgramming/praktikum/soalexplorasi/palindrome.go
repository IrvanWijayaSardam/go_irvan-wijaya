package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan sebuah kata: ")
	scanner.Scan()
	kata := scanner.Text()

	if isPalindrome(kata) {
		fmt.Println("Kata tersebut adalah palindrome.")
	} else {
		fmt.Println("Kata tersebut bukan palindrome.")
	}
}
