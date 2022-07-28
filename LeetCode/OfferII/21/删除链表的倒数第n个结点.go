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

// https://leetcode.cn/problems/SLwz0R/
// ------------------------剑指 Offer II Problem 21------------------------
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// ------------------------剑指 Offer II Problem 21------------------------
/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", removeNthFromEnd(head, n))
	}
}
