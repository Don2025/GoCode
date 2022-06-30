package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/arranging-coins/
//------------------------Leetcode Problem 441------------------------
func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		num := mid * (mid + 1)
		if num == n*2 {
			return mid
		} else if num < n*2 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

//------------------------Leetcode Problem 441------------------------
/*
 * https://leetcode.cn/problems/arranging-coins/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Println(arrangeCoins(n))
	}
}
