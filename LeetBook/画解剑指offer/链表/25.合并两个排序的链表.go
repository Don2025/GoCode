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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return head.Next
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		l1 := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		printLinkedList(l1)
		input.Scan()
		l2 := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		printLinkedList(l2)
		ans := mergeTwoLists(l1, l2)
		printLinkedList(ans)
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

func printLinkedList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		if p != head {
			fmt.Printf("->")
		}
		fmt.Printf("%d", p.Val)
	}
	println()
}

/*
 * 执行用时：4ms 在所有Go提交中击败了97.35%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了78.51%的用户
**/
