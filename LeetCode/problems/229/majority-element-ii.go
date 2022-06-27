package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/majority-element-ii/
//------------------------Leetcode Problem 229------------------------
func majorityElement(nums []int) []int {
	var ans []int
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for i, x := range cnt {
		if x > len(nums)/3 {
			ans = append(ans, i)
		}
	}
	return ans
}

//------------------------Leetcode Problem 229------------------------
/*
 * https://leetcode.cn/problems/majority-element-ii/
 * 执行用时：8ms 在所有Go提交中击败了92.86%的用户
 * 占用内存：5MB 在所有Go提交中击败了23.63%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", majorityElement(nums))
	}
}
