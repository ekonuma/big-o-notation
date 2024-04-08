package main

import "fmt"

func main() {
	sumO1(1, 2)
	sumOn([]int{1, 2, 3, 4, 5})
	sumOn2([][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}})
}

// O(1)
func sumO1(a, b int) int {
	return a + b
}

// O(n)
func sumOn(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// O(n^2)
func sumOn2(a [][]int) {
	for _, row := range a {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
