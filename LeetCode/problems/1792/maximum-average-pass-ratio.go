package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/maximum-average-pass-ratio/
//------------------------Leetcode Problem 1792------------------------
type tuple struct {
	x    float64
	a, b int
}

type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := hp{}
	for _, class := range classes {
		a, b := class[0], class[1]
		x := float64(a+1)/float64(b+1) - float64(a)/float64(b)
		heap.Push(&pq, tuple{x, a, b})
	}
	for i := 0; i < extraStudents; i++ {
		e := heap.Pop(&pq).(tuple)
		a, b := e.a+1, e.b+1
		x := float64(a+1)/float64(b+1) - float64(a)/float64(b)
		heap.Push(&pq, tuple{x, a, b})
	}
	var ans float64
	for len(pq) > 0 {
		e := heap.Pop(&pq).(tuple)
		ans += float64(e.a) / float64(e.b)
	}
	return ans / float64(len(classes))
}

//------------------------Leetcode Problem 1792------------------------
/*
 * https://leetcode.cn/problems/maximum-average-pass-ratio/
 * 执行用时：480ms 在所有Go提交中击败了38.89%的用户
 * 占用内存：33.9MB 在所有Go提交中击败了22.22%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		classes := make([][]int, n)
		for i := range classes {
			scanner.Scan()
			classes[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		ans := maxAverageRatio(classes, n)
		fmt.Printf("%f", ans)
	}
}
