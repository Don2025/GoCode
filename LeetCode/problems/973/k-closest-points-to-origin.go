package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/k-closest-points-to-origin/
//------------------------Leetcode Problem 973------------------------
type pair struct {
	Dist  int
	Point []int
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].Dist > h[j].Dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosest(points [][]int, k int) [][]int {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h)
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].Dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	var ans [][]int
	for _, p := range h {
		ans = append(ans, p.Point)
	}
	return ans
}

//------------------------Leetcode Problem 973------------------------
/*
 * https://leetcode.cn/problems/k-closest-points-to-origin/
 * 执行用时：132ms 在所有Go提交中击败了38.55%的用户
 * 占用内存：7.9MB 在所有Go提交中击败了69.28%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		points := make([][]int, n)
		for i := range points {
			scanner.Scan()
			points[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kClosest(points, k))
	}
}
