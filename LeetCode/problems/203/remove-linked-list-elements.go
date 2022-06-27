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

// https://leetcode.cn/problems/remove-linked-list-elements/
//------------------------Leetcode Problem 203------------------------
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

//------------------------Leetcode Problem 203------------------------
/*
 * https://leetcode.cn/problems/remove-linked-list-elements/
 * 执行用时：8ms 在所有Go提交中击败了66.57%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了7.78%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		ans := removeElements(head, val)
		Printf("Output: ")
		utils.PrintLinkedList(ans)
	}
}
