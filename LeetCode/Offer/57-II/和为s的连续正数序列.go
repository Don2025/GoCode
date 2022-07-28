package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// ------------------------剑指 Offer I Problem 57-II------------------------
func findContinuousSequence(target int) [][]int {
	var ans [][]int
	var sum int
	for l, r := 1, 1; r < target; r++ {
		sum += r
		for sum > target {
			sum -= l
			l++
		}
		if sum == target {
			arr := make([]int, r-l+1)
			for i := range arr {
				arr[i] = l + i
			}
			ans = append(ans, arr)
		}
	}
	return ans
}

// ------------------------剑指 Offer I Problem 57-II------------------------
/*
 * https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了90.62%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findContinuousSequence(n))
	}
}
