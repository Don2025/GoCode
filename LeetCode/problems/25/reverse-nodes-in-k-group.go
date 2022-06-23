package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
//------------------------Leetcode Problem 25------------------------
func reverseKGroup(head *ListNode, k int) *ListNode {
	var n int
	p := head
	for n < k && p != nil {
		p = p.Next
		n++
	}
	var cnt int
	var prev, next *ListNode
	cur := head
	if n == k {
		for cnt < k && cur != nil {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
			cnt++
		}
		if next != nil {
			head.Next = reverseKGroup(next, k)
		}
		return prev
	} else {
		return head
	}
}

//------------------------Leetcode Problem 25------------------------
/*
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/
 * 执行用时：4ms 在所有Go提交中击败了89.89%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了55.70%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(reverseKGroup(head, k))
	}
}
