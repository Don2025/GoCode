package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/beautiful-arrangement-ii/
//------------------------Leetcode Problem 667------------------------
func constructArray(n int, k int) []int {
	ans := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

//------------------------Leetcode Problem 667------------------------
/*
 * https://leetcode.cn/problems/beautiful-arrangement-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了33.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n, k int
		Sscanf(scanner.Text(), "%d %d", &n, &k)
		Printf("Output: %v\n", constructArray(n, k))
	}
}
