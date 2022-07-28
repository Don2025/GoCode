package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

/* Time Limit Exceeded
func goodDaysToRobBank(security []int, time int) []int {
	var ans []int
	for i := time; i < len(security)-time; i++ {
		flag := true
		for j := i-time; j < i; j++ {
			if security[j] < security[j+1] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		for j := i; j < i+time; j++ {
			if security[j] > security[j+1] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}
*/
// https://leetcode.cn/problems/find-good-days-to-rob-the-bank/
//------------------------Leetcode Problem 2100------------------------
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			right[i] = right[i+1] + 1
		}
	}
	var ans []int
	for i := 0; i < n; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

//------------------------Leetcode Problem 2100------------------------
/*
 * https://leetcode.cn/problems/find-good-days-to-rob-the-bank/
 * 执行用时：108ms 在所有Go提交中击败了18.57%的用户
 * 占用内存：10MB 在所有Go提交中击败了15.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		security := utils.StringToInts(scanner.Text())
		scanner.Scan()
		time, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", goodDaysToRobBank(security, time))
	}
}
