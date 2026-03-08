package main

import "fmt"

// speed - O(n), memory - O(n)
func containsDuplicate(nums []int) bool {
	freq := make(map[int]int, len(nums))
	for _, elem := range nums {
		freq[elem] = freq[elem] + 1
		if freq[elem] > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}
