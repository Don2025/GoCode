package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/* 提交时间：2021/09/02 21:27
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了59.74%的用户
**/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	for p := head; p != nil; p = p.Next {
		if k == 0 {
			head = head.Next
		} else {
			k--
		}
	}
	return head
}

/* 提交时间：2021/06/07 21:57
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了59.74%的用户

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var cnt int
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	cnt -= k
	ans := head
	for cnt != 0 {
		ans = ans.Next
		cnt--
	}
	return ans
}
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		head := intArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		printLinkedList(getKthFromEnd(head, k))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
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

func printLinkedList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		if p != head {
			fmt.Printf("->")
		}
		fmt.Printf("%d", p.Val)
	}
	println()
}
