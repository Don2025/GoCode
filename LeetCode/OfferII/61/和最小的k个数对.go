package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/qn8gGX/
// ------------------------剑指 Offer II Problem 61------------------------
type pair struct{ i, j int }
type hp struct {
	Data         []pair
	Nums1, Nums2 []int
}

func (h hp) Len() int { return len(h.Data) }
func (h hp) Less(i, j int) bool {
	a, b := h.Data[i], h.Data[j]
	return h.Nums1[a.i]+h.Nums2[a.j] < h.Nums1[b.i]+h.Nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }
func (h *hp) Push(v interface{}) { h.Data = append(h.Data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.Data; v := a[len(a)-1]; h.Data = a[:len(a)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.Data = append(h.Data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(h, pair{i, j + 1})
		}
	}
	return
}

// ------------------------剑指 Offer II Problem 61------------------------
/*
 * https://leetcode.cn/problems/qn8gGX/
 * 执行用时： 4ms 在所有Go提交中击败了 82.12%的用户
 * 占用内存： 3.7MB 在所有Go提交中击败了 82.12%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		Printf("Input a line of numbers separated by space:")
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kSmallestPairs(nums1, nums2, k))
		Printf("Input a line of numbers separated by space:")
	}
}
