package main

import (
	"fmt"
	"sync"
)

var (
	counter  = 0
	mutex    sync.Mutex
	wg       sync.WaitGroup
	maxCount = 100
)

func incrementCounter() {
	defer wg.Done()
	for i := 0; i < maxCount; i++ {
		mutex.Lock() // Mengunci akses ke variabel bersama
		counter++
		mutex.Unlock() // Membuka kunci setelah selesai mengakses variabel bersama
	}
}

func main() {
	wg.Add(2)

	go incrementCounter()
	go incrementCounter()

	wg.Wait()

	fmt.Printf("Nilai counter akhir: %d\n", counter)
}
