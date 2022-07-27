package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix/
//------------------------Leetcode Problem 1337------------------------
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

//------------------------Leetcode Problem 1337------------------------
/*
 * https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix/
 * 执行用时：16ms 在所有Go提交中击败了18.18%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了16.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		mat := make([][]int, n)
		for i := range mat {
			scanner.Scan()
			mat[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kWeakestRows(mat, k))
	}
}
