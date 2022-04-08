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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var ans *ListNode
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		var a, b int
		if len(s1) > 0 {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		node := &ListNode{cur, ans}
		ans = node
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		l1 := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		input.Scan()
		l2 := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		ans := addTwoNumbers(l1, l2)
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

//带值头结点
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
 * 执行用时：12ms 在所有Go提交中击败了50.69%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了53.01%的用户
**/
