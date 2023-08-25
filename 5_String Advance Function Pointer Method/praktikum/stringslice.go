package main

import "fmt"

func Mapping(slice []string) map[string]int {
	mapping := make(map[string]int)

	for _, s := range slice {
		mapping[s]++
	}

	return mapping
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
