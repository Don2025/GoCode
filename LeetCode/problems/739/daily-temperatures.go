package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/daily-temperatures/
//------------------------Leetcode Problem 739------------------------

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				ans[i] = j - i
				break
			} else if ans[j] == 0 {
				ans[i] = 0
				break
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 739------------------------
/*
 * https://leetcode.cn/problems/daily-temperatures/
 * 执行用时：620ms 在所有Go提交中击败了13.52%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了38.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tempratures := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", dailyTemperatures(tempratures))
	}
}
