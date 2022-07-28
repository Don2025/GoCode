package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/LGjMqU/
// ------------------------剑指 Offer II Problem 26------------------------
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1, l2 := head, mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
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

func mergeList(l1, l2 *ListNode) {
	var p1, p2 *ListNode
	for l1 != nil && l2 != nil {
		p1 = l1.Next
		p2 = l2.Next
		l1.Next = l2
		l1 = p1
		l2.Next = l1
		l2 = p2
	}
}

// ------------------------剑指 Offer II Problem 26------------------------
/*
 * https://leetcode.cn/problems/LGjMqU/
 * 执行用时：12ms 在所有Go提交中击败了12.71%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		reorderList(head)
		Printf("Output: ")
		utils.PrintLinkedList(head)
	}
}
