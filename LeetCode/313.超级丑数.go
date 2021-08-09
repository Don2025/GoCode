package main

import (
	"bufio"
	"container/heap"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		println(nthSuperUglyNumber(n, stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：388ms 在所有Go提交中击败了12.28%的用户
 * 占用内存：77.1MB 在所有Go提交中击败了5.26%的用户
**/
