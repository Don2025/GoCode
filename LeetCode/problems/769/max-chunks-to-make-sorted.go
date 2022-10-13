package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/max-chunks-to-make-sorted/
// ------------------------Leetcode Problem 769------------------------
func maxChunksToSorted(arr []int) int {
	var ans, max int
	for i, x := range arr {
		if x > max {
			max = x
		}
		if i == max {
			ans++
		}
	}
	return ans
}

// ------------------------Leetcode Problem 769------------------------
/*
 * https://leetcode.cn/problems/max-chunks-to-make-sorted/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了45.36%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", maxChunksToSorted(arr))
	}
}
