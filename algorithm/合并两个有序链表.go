package main


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {


	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	var tail *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	tail = head

	for {
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}

		if l1.Val > l2.Val {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		}
	}
	return head
}