package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity  O(n)
// Space Complexity O(1)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil || fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
