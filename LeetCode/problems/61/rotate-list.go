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

// https://leetcode.cn/problems/rotate-list/
//------------------------Leetcode Problem 61------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	p := head
	for p.Next != nil {
		p = p.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	p.Next = head
	for ; add > 0; add-- {
		p = p.Next
	}
	ans := p.Next
	p.Next = nil
	return ans
}

//------------------------Leetcode Problem 61------------------------
/*
 * https://leetcode.cn/problems/rotate-list/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了71.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(rotateRight(head, k))
	}
}
