package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/linked-list-components/
// ------------------------Leetcode Problem 817------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) (ans int) {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			inSet = true
			ans++
		}
	}
	return
}

// ------------------------Leetcode Problem 817------------------------
/*
 * https://leetcode.cn/problems/linked-list-components/
 * 执行用时：16ms 在所有Go提交中击败了78.57%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了32.86%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", numComponents(head, nums))
	}
}
