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

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	p := head
	for p.Next != nil {
		p = p.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	p.Next = head
	for ; add > 0; add-- {
		p = p.Next
	}
	ans := p.Next
	p.Next = nil
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		head := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := rotateRight(head, k)
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了71.09%的用户
**/
