package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
// ------------------------剑指 Offer I Problem 40------------------------

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	h := make(IntHeap, k)
	hp := &h
	copy(*hp, arr[:k+1])
	heap.Init(hp)
	for i := k; i < len(arr); i++ {
		if arr[i] < h[0] {
			heap.Pop(hp)
			heap.Push(hp, arr[i])
		}
	}
	return h
}

// ------------------------剑指 Offer I Problem 40------------------------
/*
 * https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
 * 执行用时：20ms 在所有Go提交中击败了93.64%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了35.95%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", getLeastNumbers(arr, k))
	}
}
