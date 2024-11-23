package main

import (
	"fmt"
	"sort"
)

func RemoveDuplicate(arr []int) []int {

	newMap := make(map[int]bool)
	result := []int{}

	for _, item := range arr {
		if _, found := newMap[item]; !found {
			// fmt.Print(found)
			result = append(result, item)
			newMap[item] = true
		}
	}

	fmt.Println(newMap)

	return result
}

func main() {
	arr := []int{8, 1, 2, 3, 3, 4}

	result := RemoveDuplicate(arr)
	fmt.Println(result)
	sort.Ints(result)
	fmt.Println(result)

	newMap := map[string]int{
		"apple":  10,
		"banana": 2,
	}

	if value, ok := newMap["apple"]; ok {
		fmt.Print(value)
		fmt.Println("ok")

	}
}
