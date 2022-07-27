package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/water-bottles/
//------------------------Leetcode Problem 1518------------------------
func numWaterBottles(numBottles int, numExchange int) int {
	cnt := numBottles
	for numBottles >= numExchange {
		num := numBottles / numExchange
		cnt += num
		numBottles = numBottles%numExchange + num
	}
	return cnt
}

//------------------------Leetcode Problem 1518------------------------
/*
 * https://leetcode.cn/problems/water-bottles/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var numBottles, numExchange int
		Sscanf(scanner.Text(), "%d %d", &numBottles, &numExchange)
		Printf("Output: %v\n", numWaterBottles(numBottles, numExchange))
	}
}
