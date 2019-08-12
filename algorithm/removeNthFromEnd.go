package main



type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	last := head
	for i:= 0; i< n; i++ {
		if last.Next == nil {
			if i < n-1 {
				return  head
			} else {
				// 直接删除头节点
				return head.Next
			}
		}
		last = last.Next
	}
	// first 与last同时向后移动，如果last.next==nil,那么久直接删除first节点
	// 删除
	for {
		if last.Next == nil {
			// 删除first节点
			first.Next = first.Next.Next
			break
		}
		first = first.Next
		last = last.Next
	}
	return head
}