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

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）
/* 不用栈，不用递归，直接迭代
 * 执行用时：4ms 在所有Go提交中击败了62.96%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了96.44%的用户
**/
func reversePrint(head *ListNode) []int {
	var cnt int
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	ans := make([]int, cnt)
	for p := head; p != nil; p = p.Next {
		ans[cnt-1] = p.Val
		cnt--
	}
	return ans
}

/* 递归法
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.7MB 在所有Go提交中击败了24.62%的用户
**/
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint1(head.Next), head.Val)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		head := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		printLinkedList(head)
		printIntArray(reversePrint(head))
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

func printIntArray(arr []int) {
	for _, x := range arr {
		fmt.Printf("%d ", x)
	}
	println()
}
