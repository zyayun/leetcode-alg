package main

import "fmt"

func BoubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {

	numbers := []int{8, 2, 1, 3}

	BoubbleSort(numbers)
	fmt.Println(numbers)

}
