package main

import (
	"fmt"
)

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		halfPow := pow(x, n/2)
		return halfPow * halfPow
	}
	halfPow := pow(x, (n-1)/2)
	return x * halfPow * halfPow
}

func main() {
	fmt.Println(pow(2, 3))  // Output: 8
	fmt.Println(pow(5, 3))  // Output: 125
	fmt.Println(pow(10, 2)) // Output: 100
	fmt.Println(pow(2, 5))  // Output: 32
	fmt.Println(pow(7, 3))  // Output: 343
}
