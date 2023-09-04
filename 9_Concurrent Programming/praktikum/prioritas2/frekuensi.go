package main

import (
	"fmt"
	"strings"
)

func countLetters(text string) map[rune]int {
	letterCounts := make(map[rune]int)

	for _, char := range text {
		if char != ' ' { // Abaikan spasi
			letterCounts[char]++
		}
	}

	return letterCounts
}

func main() {
	text := "Lorem Ipsum"
	letterCounts := countLetters(strings.ToLower(text)) // Ubah ke huruf kecil untuk menghindari sensitivitas huruf

	// Mencetak frekuensi kemunculan huruf
	fmt.Printf("e: %d\n", letterCounts['e'])
	fmt.Printf("i: %d\n", letterCounts['i'])
	fmt.Printf("o: %d\n", letterCounts['o'])
	fmt.Printf("t: %d\n", letterCounts['t'])
}
