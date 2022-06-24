package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/sort-colors/
//------------------------Leetcode Problem 75------------------------
func sortColors(nums []int) {
	i, j := -1, -1
	for x := range nums {
		if nums[x] < 2 {
			j++
			nums[x], nums[j] = nums[j], nums[x]
			if nums[j] < 1 {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//------------------------Leetcode Problem 75------------------------
/*
 * https://leetcode.cn/problems/sort-colors/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		sortColors(nums)
		Printf("Origin: %v\n", nums)
	}
}
