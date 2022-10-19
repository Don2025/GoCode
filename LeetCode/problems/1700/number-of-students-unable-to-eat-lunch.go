package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
//------------------------Leetcode Problem 1700------------------------
func countStudents(students []int, sandwiches []int) int {
	cnt := make([]int, 2)
	for _, student := range students {
		cnt[student]++
	}
	for _, sandwich := range sandwiches {
		if cnt[sandwich] == 0 {
			break
		}
		cnt[sandwich]--
	}
	return cnt[0] + cnt[1]
}

//------------------------Leetcode Problem 1700------------------------
/*
 * https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
 * 执行用时：4ms 在所有Go提交中击败了25.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了82.14%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		students := utils.StringToInts(scanner.Text())
		scanner.Scan()
		sandwiches := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", countStudents(students, sandwiches))
	}
}
