package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countMaxOrSubsets(nums []int) int {
	var ans, maxOr int
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if len(nums) == pos {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(countMaxOrSubsets(nums))
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
 * 执行用时：8ms 在所有Go提交中击败了65.79%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了81.58%的用户
**/
