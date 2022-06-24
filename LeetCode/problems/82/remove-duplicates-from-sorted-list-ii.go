package _2

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
//------------------------Leetcode Problem 82------------------------
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	if head.Val == p.Val {
		for p != nil && head.Val == p.Val {
			p = p.Next
		}
		head = deleteDuplicates(p)
	} else {
		head.Next = deleteDuplicates(p)
	}
	return head
}

//------------------------Leetcode Problem 82------------------------
/*
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
 * 执行用时：4ms 在所有Go提交中击败了76.14%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", deleteDuplicates(head))
	}
}
