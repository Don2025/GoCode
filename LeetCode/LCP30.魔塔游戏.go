package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func magicTower(nums []int) int {
	sum, att, ans := 1, 0, 0 // sum当前血量, att被换到最后房间的怪物伤害, ans结果
	h := new(hp)
	for _, x := range nums {
		if x < 0 {
			heap.Push(h, x)
		}
		sum += x
		for sum <= 0 {
			if len(*h) != 0 {
				ans++
				t := heap.Pop(h).(int)
				sum -= t
				att -= t
			} else {
				return -1
			}
		}
	}
	if sum > att {
		return ans
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(magicTower(nums))
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
 * 执行用时：128ms 在所有Go提交中击败了26.92%的用户
 * 占用内存：8.4MB 在所有Go提交中击败了46.15%的用户
**/
