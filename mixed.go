package main

import "fmt"

func RemoveDuplicateList(arr []int) []int {

	newMap := make(map[int]bool)

	result := []int{}

	for _, value := range arr {
		if _, ok := newMap[value]; !ok {
			result = append(result, value)
			newMap[value] = true
		}
	}

	return result
}

func BoubbleSort1(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func BinnarySearch(arr []int, target int) int {
	low, hight := 0, len(arr)-1
	for low <= hight {
		mid := (low + hight) // 2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
func main() {
	arr := []int{55, 11, 22, 22, 33, 44}
	target := 11

	result := BinnarySearch(arr, target)

	if result != -1 {
		fmt.Println("Target found index is ", result)
	} else {
		fmt.Println("Target  not found")
	}

}
