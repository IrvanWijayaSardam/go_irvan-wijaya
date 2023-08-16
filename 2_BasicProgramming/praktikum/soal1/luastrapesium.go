package main

import (
	"fmt"
)

func main() {
	var tinggi float64
	var sisiAtas float64
	var sisiBawah float64

	fmt.Print("Masukkan panjang sisi atas trapesium: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan panjang sisi bawah trapesium: ")
	fmt.Scanln(&sisiBawah)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * (sisiAtas + sisiBawah) * tinggi

	fmt.Printf("Luas trapesium: %.2f\n", luas)
}
