package main

import (
	"bufio"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/reorder-list/
//------------------------Leetcode Problem 143------------------------
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var arr []*ListNode
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p)
	}
	i, j := 0, len(arr)-1
	for i < j {
		arr[i].Next = arr[j]
		i++
		if i == j {
			break
		}
		arr[j].Next = arr[i]
		j--
	}
	arr[i].Next = nil
}

//------------------------Leetcode Problem 143------------------------
/*
 * https://leetcode.cn/problems/reorder-list/
 * 执行用时：8ms 在所有Go提交中击败了85.99%的用户
 * 占用内存：5.8MB 在所有Go提交中击败了30.34%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		reorderList(head)
		utils.PrintLinkedList(head)
	}
}
