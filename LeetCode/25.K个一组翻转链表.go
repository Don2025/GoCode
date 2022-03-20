package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var n int
	p := head
	for n < k && p != nil {
		p = p.Next
		n++
	}
	var cnt int
	var prev, next *ListNode
	cur := head
	if n == k {
		for cnt < k && cur != nil {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
			cnt++
		}
		if next != nil {
			head.Next = reverseKGroup(next, k)
		}
		return prev
	} else {
		return head
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了89.89%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了55.70%的用户
**/
