package utils

import (
	"fmt"
	"github.com/Don2025/GoCode/structures"
	"strconv"
	"strings"
)

type ByteSize float64

// Definition for order of magnitude
const (
	_           = iota             //用下划线忽略第一个变量,此时iota=0
	KB ByteSize = 1 << (10 * iota) //1 << (10*1),此时iota=1
	MB                             //1 << (10*2),此时iota=2
	GB                             //1 << (10*3),此时iota=3
	TB                             //1 << (10*4),此时iota=4
	PB                             //1 << (10*5),此时iota=5
	EB                             //1 << (10*6),此时iota=6
	ZB                             //1 << (10*7),此时iota=7
	YB                             //1 << (10*8),此时iota=8
)

// StringToInts converts string to []int.
func StringToInts(str string) (nums []int) {
	if len(str) == 0 {
		return
	}
	strs := strings.Fields(str)
	for _, ch := range strs {
		n, _ := strconv.Atoi(ch)
		nums = append(nums, n)
	}
	return
}

// IntsToString converts []int to string separated by commas.
func IntsToString(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	sb := &strings.Builder{}
	sb.WriteByte('[')
	for i := range nums {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	sb.WriteByte(']')
	return sb.String()
}

// StringToFloat64s converts string to []float64.
func StringToFloat64s(str string) (nums []float64) {
	if len(str) == 0 {
		return
	}
	strs := strings.Fields(str)
	for _, ch := range strs {
		n, _ := strconv.ParseFloat(ch, 64)
		nums = append(nums, n)
	}
	return
}

// StringToLinkList converts string to *ListNode.
func StringToLinkList(str string) *structures.ListNode {
	if len(str) == 0 {
		return nil
	}
	strs := strings.Fields(str)
	head := &structures.ListNode{}
	p := head
	for _, ch := range strs {
		n, _ := strconv.Atoi(ch)
		p.Next = &structures.ListNode{Val: n}
		p = p.Next
	}
	return head.Next
}

// IntsToLinkList converts []int to *ListNode.
func IntsToLinkList(nums []int) *structures.ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &structures.ListNode{}
	p := head
	for _, x := range nums {
		p.Next = &structures.ListNode{Val: x}
		p = p.Next
	}
	return head.Next
}

// LinkListToInts converts *ListNode to []int.
func LinkListToInts(head *structures.ListNode) (nums []int) {
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	return
}

// PrintLinkedList prints *ListNode to standard output.
func PrintLinkedList(head *structures.ListNode) {
	for p := head; p != nil; p = p.Next {
		if p != head {
			fmt.Printf("->")
		}
		fmt.Printf("%d", p.Val)
	}
	fmt.Println()
}

// StringToTreeNode converts String to *TreeNode.
func StringToTreeNode(str string) *structures.TreeNode {
	if len(str) == 0 {
		return nil
	}
	strs := strings.Fields(str)
	n := len(strs)
	if n == 0 {
		return nil
	}
	x, _ := strconv.Atoi(strs[0])
	root := &structures.TreeNode{Val: x}
	queue := make([]*structures.TreeNode, 1, n<<1)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Left = &structures.TreeNode{Val: x}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Right = &structures.TreeNode{Val: x}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
