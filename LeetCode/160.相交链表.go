package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入两个链表的头结点，返回第一个值和地址都相等的结点。
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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		linkListA := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		printLinkList(linkListA)
		input.Scan()
		linkListB := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		printLinkList(linkListB)
		//这题给的输入有问题吧，感觉输入参数少了
		// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
		ans := getIntersectionNode(linkListA, linkListB)
		if ans != nil {
			fmt.Printf("Intersected at '%d'\n", ans.Val)
		} else {
			println("null")
		}
	}
}

func intArrayToLinkList(arr []int) *ListNode {
	var head *ListNode = new(ListNode)
	p := head
	for _, x := range arr {
		p.Next = &ListNode{Val: x}
		p = p.Next
	}
	return head.Next
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

//带值头结点
func printLinkList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		if p != head {
			fmt.Printf("->")
		}
		fmt.Printf("%d", p.Val)
	}
	println()
}

/*
 * 执行用时：8ms 在所有Go提交中击败了90.61%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.71%的用户
**/
