package main

import (
	"bufio"
	"os"
	"strconv"
)

func findKthNumber(n int, k int) int {
	ans := 1
	var getSteps func(int, int) int
	getSteps = func(cur int, n int) int {
		first, last := cur, cur
		var steps int
		for first <= n {
			steps += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return steps
	}
	k--
	for k > 0 {
		steps := getSteps(ans, n)
		if steps <= k {
			k -= steps
			ans++
		} else {
			ans *= 10
			k--
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
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(findKthNumber(n, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了65.00%的用户
**/
