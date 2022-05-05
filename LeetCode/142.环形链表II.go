package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

/*
 * 执行用时：8ms 在所有Go提交中击败了26.52%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了15.36%的用户
**/
