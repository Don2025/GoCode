package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kWeakestRows(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		pow := sort.Search(len(row), func(i int) bool {
			return row[i] == 0
		})
		h = append(h, pair{pow, i})
	}
	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).idx
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		mat := make([][]int, n)
		for i := range mat {
			input.Scan()
			mat[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		fmt.Printf("%v\n", kWeakestRows(mat, k))
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
 * 执行用时：16ms 在所有Go提交中击败了18.18%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了16.78%的用户
**/
