package main

import (
	"fmt"
)

// twoSum
func twoSum(nums []int, target int) []int {
	fmt.Println("Debug: origin nums", nums)
	fmt.Println("Debug: target num", target)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}

		}
	}

	return nil
}

func twoSumHashSearch(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		if p, ok := hashMap[target-v]; ok {
			fmt.Println(p)
			return []int{p, i}
		}
		hashMap[v] = 1
		fmt.Println(hashMap)
	}
	return nil
}
func main() {
	nums := []int{5, 6, 7, 2, 9, 10}
	target := 15
	result := twoSumHashSearch(nums, target)
	fmt.Println(result)
}
