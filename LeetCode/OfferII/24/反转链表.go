package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/UHnkqh/
// ------------------------剑指 Offer II Problem 24------------------------
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ------------------------剑指 Offer II Problem 24------------------------
/*
 * https://leetcode.cn/problems/UHnkqh/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(reverseList(head))
	}
}
