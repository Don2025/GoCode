package main

type Node struct {
	Val  int
	Next *Node
}

// https://leetcode.cn/problems/4ueAj6/
// ------------------------剑指 Offer II Problem 29------------------------
func insert(head *Node, x int) *Node {
	if head == nil {
		head = &Node{x, nil}
		head.Next = head
		return head
	}
	cur := head
	for cur.Next != head {
		if cur.Next.Val < cur.Val {
			if x >= cur.Val || x <= cur.Next.Val {
				break
			}
		} else if x >= cur.Val && x <= cur.Next.Val {
			break
		}
		cur = cur.Next
	}
	cur.Next = &Node{x, cur.Next}
	return head
}

// ------------------------剑指 Offer II Problem 29------------------------
/*
 * https://leetcode.cn/problems/4ueAj6/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了74.52%的用户
**/
