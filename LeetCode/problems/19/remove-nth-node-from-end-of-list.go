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

//------------------------Leetcode Problem 19------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

//------------------------Leetcode Problem 19------------------------
/*
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了61.93%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(removeNthFromEnd(head, n))
		Printf("Input a line of numbers separated by space:")
	}
}
