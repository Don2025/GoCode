package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

//输入两个链表的头结点，返回第一个值和地址都相等的结点。
//https://leetcode.cn/problems/intersection-of-two-linked-lists/
//------------------------Leetcode Problem 160------------------------
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	you, she := headA, headB
	for you != she {
		if you != nil {
			you = you.Next
		} else {
			you = headB
		}
		if she != nil {
			she = she.Next
		} else {
			she = headA
		}
	}
	return you
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linkListA := utils.StringToListNode(scanner.Text())
		utils.PrintLinkedList(linkListA)
		scanner.Scan()
		linkListB := utils.StringToListNode(scanner.Text())
		utils.PrintLinkedList(linkListB)
		//这题给的输入有问题吧，感觉输入参数少了
		// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
		ans := getIntersectionNode(linkListA, linkListB)
		if ans != nil {
			Printf("Intersected at '%d'\n", ans.Val)
		} else {
			Println("null")
		}
	}
}

//------------------------Leetcode Problem 160------------------------
/*
 * https://leetcode.cn/problems/intersection-of-two-linked-lists/
 * 执行用时：8ms 在所有Go提交中击败了90.61%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.71%的用户
**/
