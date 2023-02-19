package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/
//------------------------Leetcode Problem 1------------------------
func countBalls(lowLimit int, highLimit int) (ans int) {
	count := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		count[sum]++
		if count[sum] > ans {
			ans = count[sum]
		}
	}
	return
}

//------------------------Leetcode Problem 1742------------------------
/*
 * https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/
 * 执行用时：24ms 在所有Go提交中击败了11.90%的用户
 * 占用内存：2MB 在所有Go提交中击败了38.09%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var lowLimit, highLimit int
		Sscanf(scanner.Text(), "%d %d", &lowLimit, &highLimit)
		Printf("Output: %v\n", countBalls(lowLimit, highLimit))
	}
}
