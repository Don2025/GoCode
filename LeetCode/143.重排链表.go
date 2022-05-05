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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var arr []*ListNode
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p)
	}
	i, j := 0, len(arr)-1
	for i < j {
		arr[i].Next = arr[j]
		i++
		if i == j {
			break
		}
		arr[j].Next = arr[i]
		j--
	}
	arr[i].Next = nil
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		head := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		reorderList(head)
		printLinkedList(head)
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
 * 执行用时：8ms 在所有Go提交中击败了85.99%的用户
 * 占用内存：5.8MB 在所有Go提交中击败了30.34%的用户
**/
