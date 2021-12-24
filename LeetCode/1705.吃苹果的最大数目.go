package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type Apple struct {
	Date int
	Num  int
}

type minHeap []Apple

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Date < h[j].Date }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(Apple)) }
func (h *minHeap) Pop() interface{} {
	ans := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ans
}

func eatenApples(apples []int, days []int) int {
	h := minHeap{}
	var cnt int
	for i := 0; i < len(days) || len(h) != 0; i++ {
		if i < len(days) && apples[i] != 0 {
			heap.Push(&h, Apple{i + days[i], apples[i]})
		}
		for len(h) != 0 && h[0].Date == i {
			heap.Pop(&h)
		}
		if len(h) != 0 && h[0].Num > 0 {
			cnt++
			h[0].Num--
			if h[0].Num == 0 {
				heap.Pop(&h)
			}
		}
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		apples := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		days := stringArrayToIntArray(strings.Fields(input.Text()))
		println(eatenApples(apples, days))
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
 * 执行用时：172ms 在所有Go提交中击败了37.50%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了100.00%的用户
**/
