package main

import (
	"bufio"
	. "fmt"
	"os"
)

func largestAltitude(gain []int) int {
	var sum, ans int
	for _, x := range gain {
		sum += x
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

//------------------------Leetcode Problem 1732------------------------
/*
 * https://leetcode.cn/problems/find-the-highest-altitude/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了72.55%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		gain := bufio.NewScanner(os.Stdin)
		Printf("Output: %v\n", largestAltitude(gain))
	}
}
