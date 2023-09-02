package main

import "fmt"

func fibonacci(number int) int {
	// Membuat memoization array untuk menyimpan hasil yang sudah dihitung
	memo := make([]int, number+1)
	return fibonacciRecursive(number, memo)
}

func fibonacciRecursive(number int, memo []int) int {
	// Basis case: Fibonacci dari 0 dan 1 adalah 0 dan 1
	if number <= 1 {
		return number
	}

	// Jika hasil sudah dihitung sebelumnya, gunakan hasil tersebut
	if memo[number] != 0 {
		return memo[number]
	}

	// Hitung Fibonacci secara rekursif dan simpan hasilnya di memoization array
	memo[number] = fibonacciRecursive(number-1, memo) + fibonacciRecursive(number-2, memo)
	return memo[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
