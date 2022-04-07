package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	var ans int
	for l <= r {
		mid := l + (r-l)>>1
		var cnt int
		for i := 1; i <= m; i++ {
			cnt += min(mid/i, n)
		}
		if cnt >= k {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		m, _ := strconv.Atoi(arr[0])
		n, _ := strconv.Atoi(arr[1])
		k, _ := strconv.Atoi(arr[2])
		println(findKthNumber(m, n, k))
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了37.50%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了43.75%的用户
**/
