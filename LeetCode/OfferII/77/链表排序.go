package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/7WHec2/
// ------------------------剑指 Offer II Problem 77------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	t, t1, t2 := dummyHead, head1, head2
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			t.Next = t1
			t1 = t1.Next
		} else {
			t.Next = t2
			t2 = t2.Next
		}
		t = t.Next
	}
	if t1 != nil {
		t.Next = t1
	} else if t2 != nil {
		t.Next = t2
	}
	return dummyHead.Next
}

// ------------------------剑指 Offer II Problem 77------------------------
/*
 * https://leetcode.cn/problems/7WHec2/
 * 执行用时： 28ms 在所有Go提交中击败了 41.42%的用户
 * 占用内存： 7.1MB 在所有Go提交中击败了 94.97%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		head := utils.StringToLinkList(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(sortList(head))
		Printf("Input a line of numbers separated by space:")
	}
}
