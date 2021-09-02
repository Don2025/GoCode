package main

import (
	"bufio"
	"os"
)

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/* 提交时间：2021/09/02 21:27
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了59.74%的用户
**/
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

	}
}
