package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan Nomor: ")
	fmt.Scanln(&bilangan)

	if bilangan%2 == 0 {
		fmt.Printf("%d adalah bilangan genap.\n", bilangan)
	} else {
		fmt.Printf("%d adalah bilangan ganjil.\n", bilangan)
	}
}
