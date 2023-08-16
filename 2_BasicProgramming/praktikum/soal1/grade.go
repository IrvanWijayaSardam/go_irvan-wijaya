package main

import (
	"fmt"
)

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	if nilai >= 0 && nilai <= 100 {
		var grade string
		switch {
		case nilai >= 80:
			grade = "A"
		case nilai >= 65:
			grade = "B"
		case nilai >= 50:
			grade = "C"
		case nilai >= 35:
			grade = "D"
		default:
			grade = "E"
		}
		fmt.Printf("Grade: %s\n", grade)
	} else {
		fmt.Println("Nilai Invalid")
	}
}
