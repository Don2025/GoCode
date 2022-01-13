package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

/*
 * 执行用时：4ms 在所有Go提交中击败了78.32%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了60.51%的用户
**/
