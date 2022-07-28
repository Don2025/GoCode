package main

type Node struct {
	Val          int
	Next, Random *Node
}

// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
// ------------------------剑指 Offer I Problem 35------------------------
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{Val: p.Val, Next: p.Next}
	}
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	ans := head.Next
	for p := head; p != nil; p = p.Next {
		np := p.Next
		p.Next = p.Next.Next
		if np.Next != nil {
			np.Next = np.Next.Next
		}
	}
	return ans
}

// ------------------------剑指 Offer I Problem 35------------------------
/*
 * https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了75.36%的用户
**/
