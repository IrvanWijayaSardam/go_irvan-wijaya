package main

import (
	"fmt"
	"strconv"
)

func generateBinaryArray(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		binaryStr := strconv.FormatInt(int64(i), 2)
		ans[i] = binaryStr
	}
	return ans
}

func main() {
	n := 3
	result := generateBinaryArray(n)
	fmt.Println(result)
}
