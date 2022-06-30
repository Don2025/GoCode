package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/teemo-attacking/
//------------------------Leetcode Problem 495------------------------
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := duration
	for i := len(timeSeries) - 2; i >= 0; i-- {
		if timeSeries[i]+duration < timeSeries[i+1] {
			sum += duration
		} else {
			sum += timeSeries[i+1] - timeSeries[i]
		}
	}
	return sum
}

//------------------------Leetcode Problem 495------------------------
/*
 * https://leetcode.cn/problems/teemo-attacking/
 * 执行用时：44ms 在所有Go提交中击败了62.31%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了93.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		timeSeries := utils.StringToInts(scanner.Text())
		scanner.Scan()
		duration, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findPoisonedDuration(timeSeries, duration))
	}
}
