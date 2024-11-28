package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val: val,
	}

}
// 二叉树遍历 前序
func PreOderTree(head *TreeNode) []int {
	if head == nil {
		return []int{}
	}

	curr := head
	result := []int{}

	left := PreOderTree(curr.left)
	right := PreOderTree(curr.right)

	result = append(result, left...)
	result = append(result, right...)
	result = append(result, curr.val)

	return result

}

func main() {
	root := NewTreeNode(1)
	root.left = NewTreeNode(2)
	root.left.left = NewTreeNode(4)
	root.left.right = NewTreeNode(5)

	root.right = NewTreeNode(3)

	result := PreOderTree(root)

	fmt.Println(result)
}
