package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time/
//------------------------Leetcode Problem 1450------------------------
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var cnt int
	for i, s := range startTime {
		if s <= queryTime && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}

//------------------------Leetcode Problem 1450------------------------
/*
 * https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了68.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		startTime := utils.StringToInts(scanner.Text())
		scanner.Scan()
		endTime := utils.StringToInts(scanner.Text())
		scanner.Scan()
		queryTime, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", busyStudent(startTime, endTime, queryTime))
	}
}
