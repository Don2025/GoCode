package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/xx4gT2/
// ------------------------剑指 Offer II Problem 76------------------------

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *intHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func findKthLargest(nums []int, k int) int {
	hp := &intHeap{}
	for _, x := range nums {
		heap.Push(hp, x)
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	return (*hp)[0]
}

// ------------------------剑指 Offer II Problem 76------------------------
/*
 * https://leetcode.cn/problems/xx4gT2/
 * 执行用时：4ms 在所有Go提交中击败了91.57%的用户
 * 占用内存：4MB 在所有Go提交中击败了20.81%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findKthLargest(nums, k))
		Printf("Input a line of numbers separated by space:")
	}
}
