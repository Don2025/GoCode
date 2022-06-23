package _

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/add-two-numbers
//------------------------Leetcode Problem 2------------------------
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, head *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

//------------------------Leetcode Problem 2------------------------
/*
 * https://leetcode.cn/problems/add-two-numbers
 * 执行用时：12ms 在所有Go提交中击败了48.74%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了72.79%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l1 := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		l2 := utils.StringToListNode(scanner.Text())
		Printf("%v\n", addTwoNumbers(l1, l2))
	}
}
