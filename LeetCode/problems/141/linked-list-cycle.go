package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/linked-list-cycle/
//------------------------Leetcode Problem 141------------------------
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//------------------------Leetcode Problem 141------------------------
/*
 * https://leetcode.cn/problems/linked-list-cycle/
 * 执行用时：8ms 在所有Go提交中击败了76.51%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了64.45%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", hasCycle(head))
	}
}
