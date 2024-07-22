/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var curr *ListNode
	for i := 0; i < len(lists); i++ {
		curr = mergeTwoLists(curr, lists[i])
	}
	return curr
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Nextfunc merge(left, right, head *ListNode) *ListNode {
    current := head
    var tail *ListNode
    for left != nil && right != nil {
        if left.Value < right.Value {
            current.Next = left
            left = left.Next
        } else {
            current.Next = right
            right = right.Next
        }
        current = current.Next
    }

    if left != nil {
        current.Next = left
    } else {
        current.Next = right
    }

    tail = current
    while tail.Next != nil {
        tail = tail.Next
    }

    return tail
}
	} else {
		head = list2
		list2 = list2.Next
	}

	current := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return head
}
