package main

import "fmt"

func BinnarySearch(arr []int, target int) int {
	low, hight := 0, len(arr)-1

	for low <= hight {
		mid := (low + hight) // 2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1
		} else {
			hight = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 6

	result := BinnarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Debug: found target index:%d\n", result)
	} else {
		fmt.Printf("Debug: target index is not found:%d\n", result)
	}

}
