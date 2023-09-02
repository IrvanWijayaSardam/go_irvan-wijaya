package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 1 {
		return 0 // Tidak ada biaya jika hanya ada satu batu atau tidak ada batu
	}

	// Inisialisasi slice untuk menyimpan biaya minimum
	cost := make([]int, n)
	cost[0] = 0
	cost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	// Menghitung biaya minimum
	for i := 2; i < n; i++ {
		option1 := cost[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		option2 := cost[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		cost[i] = int(math.Min(float64(option1), float64(option2)))
	}

	// Biaya minimum total untuk mencapai batu terakhir
	return cost[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))             // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))     // 40
	fmt.Println(Frog([]int{10, 20, 10, 30, 50, 10, 60})) // 70
}
