// https://leetcode.com/problems/product-of-array-except-self/description/
package main

import "fmt"

// speed - O(n), memory - O(n)
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	prevLeft := 1
	prevRight := 1
	for i, el := range nums {
		prevLeft = prevLeft * el
		left[i] = prevLeft
		prevRight = prevRight * nums[len(nums)-1-i]
		right[len(nums)-1-i] = prevRight
	}
	nums[0] = right[1]
	nums[len(nums)-1] = left[len(nums)-2]
	for j := 1; j < len(nums)-1; j++ {
		nums[j] = left[j-1] * right[j+1]
	}
	return nums
}

// speed - O(n), memory - O(1)
func productExceptSelfOptimize(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	for j := 1; j < len(nums); j++ {
		left[j] = left[j-1] * nums[j-1]
	}

	suffix := 1
	for k := len(nums) - 1; k > -1; k-- {
		left[k] = left[k] * suffix
		suffix = suffix * nums[k]
	}
	return left
}

func main() {
	fmt.Println(productExceptSelfOptimize([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
