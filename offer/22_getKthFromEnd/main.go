package main

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if k > 0 {
			k--
			fast = fast.Next
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}
