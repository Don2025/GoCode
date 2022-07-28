package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/aMhZSa/
// ------------------------剑指 Offer II Problem 27------------------------
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := middleNode(head)                  // 找到前半部分链表尾结点
	secondHalfStart := reverseList(firstHalfEnd.Next) // 反转后半部分链表
	ans := true
	for p1, p2 := head, secondHalfStart; ans && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}
	firstHalfEnd.Next = reverseList(secondHalfStart) // 还原链表
	return ans
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ------------------------剑指 Offer II Problem 27------------------------
/*
 * https://leetcode.cn/problems/aMhZSa/
 * 执行用时：108ms 在所有Go提交中击败了99.05%的用户
 * 占用内存：10MB 在所有Go提交中击败了70.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", isPalindrome(head))
	}
}
