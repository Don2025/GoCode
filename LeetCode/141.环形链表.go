package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/*
 * 执行用时：8ms 在所有Go提交中击败了76.51%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了64.45%的用户
**/
