package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/
//------------------------Leetcode Problem 793------------------------
func preimageSizeFZF(k int) int {
	zeta := func(n int) (ans int) {
		for n > 0 {
			n /= 5
			ans += n
		}
		return
	}
	nx := func(n int) (ans int) {
		return sort.Search(5*n, func(x int) bool {
			return zeta(x) >= n
		})
	}
	return nx(k+1) - nx(k)
}

//------------------------Leetcode Problem 793------------------------
/*
 * https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了74.58%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", preimageSizeFZF(k))
	}
}
