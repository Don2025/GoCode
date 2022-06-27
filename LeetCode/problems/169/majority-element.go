package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/majority-element/
//------------------------Leetcode Problem 169------------------------
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

//------------------------Leetcode Problem 169------------------------
/*
 * https://leetcode.cn/problems/majority-element/
 * 执行用时：12ms 在所有Go提交中击败了94.07%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了51.98%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", majorityElement(nums))
	}
}
