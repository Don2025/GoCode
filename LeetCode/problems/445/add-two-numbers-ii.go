package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/add-two-numbers-ii/
// ------------------------Leetcode Problem 445------------------------
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var ans *ListNode
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		var a, b int
		if len(s1) > 0 {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		node := &ListNode{cur, ans}
		ans = node
	}
	return ans
}

// ------------------------Leetcode Problem 445------------------------
/*
 * https://leetcode.cn/problems/add-two-numbers-ii/
 * 执行用时：12ms 在所有Go提交中击败了50.69%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了53.01%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l1 := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		l2 := utils.StringToListNode(scanner.Text())
		ans := addTwoNumbers(l1, l2)
		Printf("Output: ")
		utils.PrintLinkedList(ans)
	}
}
