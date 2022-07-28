package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/XagZNi/
// ------------------------剑指 Offer II Problem 37------------------------
func asteroidCollision(asteroids []int) []int {
	var ans []int
	for _, a := range asteroids {
		for a < 0 && len(ans) > 0 && ans[len(ans)-1] > 0 {
			sum := ans[len(ans)-1] + a
			if sum >= 0 {
				a = 0
			}
			if sum <= 0 {
				ans = ans[:len(ans)-1]
			}
		}
		if a != 0 {
			ans = append(ans, a)
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 37------------------------
/*
 * https://leetcode.cn/problems/XagZNi/
 * 执行用时：12ms 在所有Go提交中击败了20.16%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了51.94%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		asteroids := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", asteroidCollision(asteroids))
	}
}
