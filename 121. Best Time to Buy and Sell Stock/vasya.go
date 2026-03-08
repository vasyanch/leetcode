package main

import "fmt"

// speed - O(n^2), memory - O(1)
func maxProfitFirst(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > result {
				result = diff
			}
		}
	}
	return result
}

// speed - O(n), memory - O(1)
func maxProfit(prices []int) int {
	result := 0
	minEl := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minEl {
			minEl = prices[i]
		}
		diff := prices[i] - minEl
		if diff > result {
			result = diff
		}
	}
	return result
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{2, 4, 1}))
}
