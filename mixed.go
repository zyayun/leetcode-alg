package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr.Next = next
	}

	return prev
}
func main() {
	root1 := &ListNode{Val: 1}
	root2 := &ListNode{Val: 2}
	root3 := &ListNode{Val: 3}
	root4 := &ListNode{Val: 4}
	root5 := &ListNode{Val: 5}

	root1.Next = root2
	root2.Next = root3
	root3.Next = root4
	root4.Next = root5

	result := ReverseList(root1)

	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("->%d\t", curr.Val)
	}
}
