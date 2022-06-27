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

// https://leetcode.cn/problems/the-skyline-problem/
// --------------------------Leetcode Problem 218--------------------------
type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	boundaries := make([]int, 0, n<<1)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)
	var ans [][]int
	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}
		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return ans
}

// --------------------------Leetcode Problem 218--------------------------
/*
 * https://leetcode.cn/problems/the-skyline-problem/
 * 执行用时：8ms 在所有Go提交中击败了97.98%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了73.74%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		buildings := make([][]int, n)
		for i := range buildings {
			scanner.Scan()
			buildings[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", getSkyline(buildings))
	}
}
