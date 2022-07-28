package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

/* 提交时间：2021/06/07 21:57
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了59.74%的用户

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var cnt int
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	cnt -= k
	ans := head
	for cnt != 0 {
		ans = ans.Next
		cnt--
	}
	return ans
}
**/
// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
// ------------------------剑指 Offer I Problem 22------------------------

// ListNode is Definition for singly-linked list.
type ListNode = structures.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	for p := head; p != nil; p = p.Next {
		if k == 0 {
			head = head.Next
		} else {
			k--
		}
	}
	return head
}

// ------------------------剑指 Offer I Problem 22------------------------
/*
 * https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了59.74%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", getKthFromEnd(head, k))
	}
}
