package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
	"strings"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head.Next
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		l1 := utils.IntArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
		input.Scan()
		l2 := utils.IntArrayToLinkList(stringArrayToIntArray(strings.Fields(input.Text())))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了73.83%的用户
**/
