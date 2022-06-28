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

// https://leetcode.cn/problems/super-ugly-number/
//------------------------Leetcode Problem 313------------------------
//最小堆
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthSuperUglyNumber(n int, primes []int) (ugly int) {
	seen := map[int]bool{1: true}
	h := &hp{[]int{1}} //将最小的超级丑数1加入堆中
	for i := 0; i < n; i++ {
		ugly = heap.Pop(h).(int) //每次取出的堆顶元素就是堆中最小的超级丑数
		for _, prime := range primes {
			//将数组primes中的每个质数和x的乘积分别加入堆
			if next := ugly * prime; !seen[next] {
				seen[next] = true
				heap.Push(h, next)
			}
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		primes := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", nthSuperUglyNumber(n, primes))
	}
}

//------------------------Leetcode Problem 313------------------------
/*
 * https://leetcode.cn/problems/super-ugly-number/
 * 执行用时：388ms 在所有Go提交中击败了12.28%的用户
 * 占用内存：77.1MB 在所有Go提交中击败了5.26%的用户
**/
