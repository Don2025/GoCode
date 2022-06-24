package _3

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

//https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
//------------------------Leetcode Problem 83------------------------
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//------------------------Leetcode Problem 83------------------------
/*
 * //https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
 * 执行用时：4ms 在所有Go提交中击败了78.32%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了60.51%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", deleteDuplicates(head))
	}
}
