package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/asteroid-collision/
//------------------------Leetcode Problem 735------------------------
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

//------------------------Leetcode Problem 735------------------------
/*
 * https://leetcode.cn/problems/asteroid-collision/
 * 执行用时：8ms 在所有Go提交中击败了85.00%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了93.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		asteroids := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", asteroidCollision(asteroids))
	}

}
