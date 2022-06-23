package main

import (
	"bufio"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/merge-k-sorted-lists/
//------------------------Leetcode Problem 23------------------------
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var merge2Lists func(*ListNode, *ListNode) *ListNode
	merge2Lists = func(l1 *ListNode, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		var head *ListNode
		if l1.Val <= l2.Val {
			head = l1
			head.Next = merge2Lists(l1.Next, l2)
		} else {
			head = l2
			head.Next = merge2Lists(l1, l2.Next)
		}
		return head
	}
	if len(lists) == 2 {
		return merge2Lists(lists[0], lists[1])
	}
	mid := len(lists) >> 1
	l1 := make([]*ListNode, mid)
	for i := 0; i < mid; i++ {
		l1[i] = lists[i]
	}
	l2 := make([]*ListNode, len(lists)-mid)
	for i, j := mid, 0; i < len(lists); i++ {
		l2[j] = lists[i]
		j++
	}
	return merge2Lists(mergeKLists(l1), mergeKLists(l2))
}

//------------------------Leetcode Problem 23------------------------
/*
 * https://leetcode.cn/problems/merge-k-sorted-lists/
 * 执行用时：8ms 在所有Go提交中击败了89.22%的用户
 * 占用内存：7MB 在所有Go提交中击败了5.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lists := make([]*ListNode, n)
		for i := range lists {
			scanner.Scan()
			lists[i] = utils.StringToListNode(scanner.Text())
		}
		utils.PrintLinkedList(mergeKLists(lists))
	}
}
