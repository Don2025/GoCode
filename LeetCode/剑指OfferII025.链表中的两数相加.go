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
	var nums1, nums2 []int
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var ans *ListNode
	for len(nums1) > 0 || len(nums2) > 0 || carry != 0 {
		var n1, n2 int
		if len(nums1) > 0 {
			n1 = nums1[len(nums1)-1]
			nums1 = nums1[:len(nums1)-1]
		}
		if len(nums2) > 0 {
			n2 = nums2[len(nums2)-1]
			nums2 = nums2[:len(nums2)-1]
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum %= 10
		ans = &ListNode{sum, ans}
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
 * 执行用时：20ms 在所有Go提交中击败了12.01%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了48.01%的用户
**/
