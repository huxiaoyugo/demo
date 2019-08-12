package main

import "fmt"

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */


 func main() {
 	defer fmt.Println(1)
	 defer fmt.Println(2)
	 defer fmt.Println(3)
	return
 }
//func reverseKGroup(head *ListNode, k int) *ListNode {
//
//	if k <= 1 {
//		return head
//	}
//
//	var h, tail,pNode, kNode, tt *ListNode
//	var t = new(ListNode)
//
//	nextHead := head
//
//	for j:=k; j<= k; j++ {
//
//		if j == k {
//			j=0
//			if h == nil {
//				h = t.Next
//			} else {
//				tail.Next = t.Next
//			}
//			t.Next = nil
//			tail = tt
//
//			if nextHead == nil {
//				if h != nil {
//					return h
//				}
//				return head
//			}
//
//			// 跟新新的
//			pNode = nextHead
//			kNode = nextHead
//			// 向前搜索k个
//			for i:=0;i<k-1;i++ {
//				kNode = kNode.Next
//				if kNode == nil {
//					if h != nil {
//						tail.Next = pNode
//						return h
//					}
//					return head
//				}
//			}
//			nextHead = kNode.Next
//		}
//
//		// pNode,kNode进行掉头，遍历pNode想t中的第一个节点后插入
//		if t.Next == nil {
//			tt = pNode
//		}
//		tmp := pNode.Next
//		pNode.Next = t.Next
//		t.Next = pNode
//		pNode = tmp
//	}
//	return h
//}