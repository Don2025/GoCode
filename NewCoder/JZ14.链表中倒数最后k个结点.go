package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil || k <= 0 {
		return nil
	}
	prev := &ListNode{0, pHead}
	cur := pHead
	for i := 0; i < k; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		prev = prev.Next
	}
	return prev.Next
}

/*
 * 运行时间：11ms 超过2.96%用Go提交的代码
 * 占用内存：796KB 超过95.32%用Go提交的代码
**/
