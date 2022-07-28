package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/lMSNwu/
//--------------------------剑指 Offer II Problem 25------------------------
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nums1, nums2 []int
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var ans *ListNode
	for len(nums1) > 0 || len(nums2) > 0 || carry != 0 {
		var n1, n2 int
		if len(nums1) > 0 {
			n1 = nums1[len(nums1)-1]
			nums1 = nums1[:len(nums1)-1]
		}
		if len(nums2) > 0 {
			n2 = nums2[len(nums2)-1]
			nums2 = nums2[:len(nums2)-1]
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum %= 10
		ans = &ListNode{sum, ans}
	}
	return ans
}

//--------------------------剑指 Offer II Problem 25------------------------
/*
 * https://leetcode.cn/problems/lMSNwu/
 * 执行用时：20ms 在所有Go提交中击败了12.01%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了48.01%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l1 := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		l2 := utils.StringToListNode(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(addTwoNumbers(l1, l2))
	}
}
