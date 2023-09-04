package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		fmt.Printf("%d ", x*i)
		time.Sleep(3 * time.Second) // Menunggu 3 detik sebelum mencetak angka berikutnya
	}
}

func main() {
	x := 5 // Ganti dengan kelipatan yang ingin Anda cetak
	go printMultiples(x)

	time.Sleep(30 * time.Second)
}
