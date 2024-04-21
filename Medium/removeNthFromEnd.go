package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	elem := dummy

	for i := 0; i <= n; i++ {
		elem = elem.Next
	}

	for elem != nil {
		curr = curr.Next
		elem = elem.Next
	}

	curr.Next = curr.Next.Next
	return dummy.Next
}
