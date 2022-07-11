package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/
//------------------------Leetcode Problem 1217------------------------
func minCostToMoveChips(position []int) int {
	var odd, even int
	for i := 0; i < len(position); i++ {
		if position[i]&1 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1217------------------------
/*
 * https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		position := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minCostToMoveChips(position))
	}
}
