package main

import "fmt"

func PairSum(arr []int, target int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil // Jika tidak ditemukan pasangan
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0 1]
}
