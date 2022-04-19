package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := getLeastNumbers(arr, k)
		fmt.Printf("%v\n", ans)
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

/*
 * 执行用时：20ms 在所有Go提交中击败了93.64%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了35.95%的用户
**/
