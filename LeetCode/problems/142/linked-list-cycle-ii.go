package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/linked-list-cycle-ii/
//------------------------Leetcode Problem 142------------------------
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

//------------------------Leetcode Problem 142------------------------
/*
 * https://leetcode.cn/problems/linked-list-cycle-ii/
 * 执行用时：8ms 在所有Go提交中击败了26.52%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了15.36%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", detectCycle(head))
	}
}
