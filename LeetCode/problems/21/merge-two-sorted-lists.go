package main

import (
	"bufio"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/merge-two-sorted-lists/
//------------------------Leetcode Problem 21------------------------
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head.Next
}

//------------------------Leetcode Problem 21------------------------
/*
 * https://leetcode.cn/problems/merge-two-sorted-lists/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了73.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l1 := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		l2 := utils.StringToListNode(scanner.Text())
		utils.PrintLinkedList(mergeTwoLists(l1, l2))
	}
}
