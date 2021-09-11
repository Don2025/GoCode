package main

import (
	"bufio"
	"os"
	"strconv"
)

/* TLE
func findIntegers(n int) int {
	var cnt int
	for i := 0; i <= n; i++ {
		if i & (i>>1) == 0 {
			cnt++;
		}
	}
	return cnt;
}
*/

func findIntegers(n int) int {
	dp := [31]int{1, 1}
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	var cnt int //记录n的二进制位数
	if n < 2 {
		return n + 1
	}
	for n>>cnt != 0 {
		cnt++
	}
	if n>>(cnt-2) == 3 {
		return dp[cnt]
	} else {
		mask := (1 << (cnt - 1)) - 1 //生成cnt-2个1
		return dp[cnt-1] + findIntegers(n&mask)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(findIntegers(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了16.67%的用户
**/
