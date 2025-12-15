// Find union of two slices.
package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{3, 4, 5, 6}
	unionSl := union(sl1, sl2)
	fmt.Println("Union of sl1 and sl2 is ---", unionSl)
}

func union(a, b []int) []int {
	m := make(map[int]bool)

	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		m[v] = true
	}

	// Result slice

	result := []int{}

	// Append the unique elements from the map

	for k := range m {
		result = append(result, k)
	}

	return result
}
