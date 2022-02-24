package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	if head.Val == p.Val {
		for p != nil && head.Val == p.Val {
			p = p.Next
		}
		head = deleteDuplicates(p)
	} else {
		head.Next = deleteDuplicates(p)
	}
	return head
}

/*
 * 执行用时：4ms 在所有Go提交中击败了76.14%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了100.00%的用户
**/
