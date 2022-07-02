package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
//------------------------Leetcode Problem 668------------------------
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

//------------------------Leetcode Problem 668------------------------
/*
 * https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
 * 执行用时：16ms 在所有Go提交中击败了37.50%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了43.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n, k int
		_, err := Sscanf(scanner.Text(), "%d %d %d", &m, &n, &k)
		if err != nil {
			break
		}
		Printf("Output: %d\n", findKthNumber(m, n, k))
	}
}
