package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/distance-between-bus-stops/
//------------------------Leetcode Problem 1184------------------------
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	var sum1, sum2 int
	for i, d := range distance {
		if i >= start && i < destination {
			sum1 += d
		} else {
			sum2 += d
		}
	}
	return min(sum1, sum2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1184------------------------
/*
 * https://leetcode.cn/problems/distance-between-bus-stops/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了65.22%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		distance := utils.StringToInts(scanner.Text())
		scanner.Scan()
		var start, destination int
		Sscanf(scanner.Text(), "%d %d", &start, &destination)
		Printf("Output: %v\n", distanceBetweenBusStops(distance, start, destination))
	}
}
