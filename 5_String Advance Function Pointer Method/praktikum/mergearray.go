package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make([]string, 0, len(arrayA)+len(arrayB))
	seenNames := make(map[string]bool)

	for _, name := range arrayA {
		if !seenNames[name] {
			merged = append(merged, name)
			seenNames[name] = true
		}
	}

	for _, name := range arrayB {
		if !seenNames[name] {
			merged = append(merged, name)
			seenNames[name] = true
		}
	}

	return merged
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
