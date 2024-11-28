package main

import "fmt"

type BinnaryTreev1 struct {
	Val   int
	Left  *BinnaryTreev1
	Right *BinnaryTreev1
}

func PreOderTreev1(root *BinnaryTreev1) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	left := PreOderTreev1(root.Left)
	right := PreOderTreev1(root.Right)

	// result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, root.Val)
	result = append(result, right...)

	return result
}

func main() {
	root := &BinnaryTreev1{Val: 1}
	root.Left = &BinnaryTreev1{Val: 2}
	root.Left.Left = &BinnaryTreev1{Val: 3}
	root.Left.Right = &BinnaryTreev1{Val: 4}
	root.Right = &BinnaryTreev1{Val: 5}

	result := PreOderTreev1(root)

	fmt.Print(result)

}
