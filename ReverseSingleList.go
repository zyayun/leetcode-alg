package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	for curr := node1; curr != nil; curr = curr.Next {
		fmt.Printf("->%d \t", curr.Val)
	}

	fmt.Println()
	result := Reverse(node1)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("->%d \t", curr.Val)
	}
}

// // v2 增加一些辅助方法
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func Reverse(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var prev *ListNode = nil
// 	curr := head
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	return prev
// }

// func Insert(head *ListNode, val int) {
// 	if head == nil {
// 		return
// 	}
// 	newNode := &ListNode{Val: val}

// 	curr := head
// 	for curr.Next != nil {
// 		curr = curr.Next
// 	}
// 	curr.Next = newNode
// }

// func PrintListNode(head *ListNode) {
// 	if head == nil {
// 		return
// 	}

// 	curr := head

// 	for curr != nil {
// 		fmt.Printf("-> %d\t", curr.Val)
// 		curr = curr.Next
// 	}
// 	fmt.Println()
// }

// func main() {

// 	root := &ListNode{Val: 1}
// 	Insert(root, 2)
// 	Insert(root, 3)
// 	Insert(root, 4)
// 	Insert(root, 5)

// 	PrintListNode(root)

// 	result := Reverse(root)
// 	PrintListNode(result)

// }
