package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/height-checker/
//------------------------Leetcode Problem 1051------------------------
func heightChecker(heights []int) int {
	sorted := append([]int(nil), heights...)
	sort.Ints(sorted)
	var ans int
	for i := range heights {
		if heights[i] != sorted[i] {
			ans++
		}
	}
	return ans
}

//------------------------Leetcode Problem 1051------------------------
/*
 * https://leetcode.cn/problems/height-checker/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了46.11%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		heights := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", heightChecker(heights))
		Printf("Input a line of numbers separated by space:")
	}
}
