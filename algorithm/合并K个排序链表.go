package main

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	var head, tail *ListNode
	for {
		minIndex :=findMin(&lists)
		if minIndex == -1 {
			break
		}
		if lists[minIndex] == nil {
			break
		}
		if head == nil {
			head = lists[minIndex]
			tail = head
		} else {
			tail.Next = lists[minIndex]
			tail = tail.Next
		}
		lists[minIndex] = lists[minIndex].Next
	}
	return head
}

func findMin(lists *[]*ListNode) int {
	if len(*lists) == 0 {
		return -1
	}
	minIndex := 0
	validCount := 0
	for index, it := range *lists {
		if it != nil {
			if (*lists)[minIndex] == nil || it.Val < (*lists)[minIndex].Val{
				minIndex = validCount
			}
			if validCount < index {
				(*lists)[validCount] = it
			}
			validCount++
		}
	}
	*lists = (*lists)[0:validCount]

	if len(*lists) == 0 {
		return -1
	}
	return minIndex
}
