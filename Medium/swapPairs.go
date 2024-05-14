/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head.Next.Next
	newHead := head.Next
	newHead.Next = head
	head.Next = swapPairs(curr)
	return newHead
}