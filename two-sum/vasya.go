package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		dif := target - v
		if secondIndex, ok := numsMap[dif]; ok {
			result[0] = i
			result[1] = secondIndex
			break
		}
		numsMap[v] = i
	}
	return result
}
