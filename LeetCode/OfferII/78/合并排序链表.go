package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/vvXgSW/
// ------------------------剑指 Offer II Problem 78------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type listHeap []*ListNode

func (h listHeap) Len() int            { return len(h) }
func (h listHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h listHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *listHeap) Push(v interface{}) { *h = append(*h, v.(*ListNode)) }
func (h *listHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	h := new(listHeap)
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	head := new(ListNode)
	cur := head
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		node = node.Next
		if node != nil {
			heap.Push(h, node)
		}
		cur = cur.Next
	}
	return head.Next
}

// ------------------------剑指 Offer II Problem 78------------------------
/*
 * https://leetcode.cn/problems/vvXgSW/
 * 执行用时： 8ms 在所有Go提交中击败了 89.22%的用户
 * 占用内存： 5.5MB 在所有Go提交中击败了 25.49%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a number that means size of lists:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lists := make([]*ListNode, n)
		for i := range lists {
			Printf("Input a line of numbers separated by space that means list%d:", i+1)
			scanner.Scan()
			lists[i] = utils.StringToLinkList(scanner.Text())
		}
		Printf("Output: ")
		utils.PrintLinkedList(mergeKLists(lists))
		Printf("Input a number that means size of lists:")
	}
}
