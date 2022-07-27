package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
//------------------------Leetcode Problem 1290------------------------
func getDecimalValue(head *ListNode) int {
	var ans int
	for head != nil {
		ans <<= 1
		ans |= head.Val
		head = head.Next
	}
	return ans
}

//------------------------Leetcode Problem 1290------------------------
/*
 * https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
 * 执行用时：8ms 在所有Go提交中击败了90.61%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", getDecimalValue(head))
	}
}
