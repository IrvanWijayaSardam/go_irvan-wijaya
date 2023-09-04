package main

import (
	"fmt"
	"sync"
)

func printMultiplesOfThree(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 3; ; i += 3 {
		ch <- i // Mengirim kelipatan 3 ke dalam channel
	}
}

func main() {
	ch := make(chan int, 5) // Buat channel dengan buffer sebanyak 5 elemen
	var wg sync.WaitGroup

	wg.Add(1)
	go printMultiplesOfThree(ch, &wg)

	for i := 0; i < 5; i++ {
		multiple := <-ch // Menerima nilai kelipatan 3 dari channel
		fmt.Printf("Kelipatan 3 ke-%d: %d\n", i+1, multiple)
	}

	close(ch) // Menutup channel setelah selesai menggunakannya
	wg.Wait() // Menunggu goroutine selesai
}
