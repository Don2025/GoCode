package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/single-number-iii/
//------------------------Leetcode Problem 260------------------------
func singleNumber(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	var ans []int
	for key, val := range cnt {
		if val == 1 {
			ans = append(ans, key)
		}
	}
	return ans
}

//------------------------Leetcode Problem 260------------------------
/*
 * https://leetcode.cn/problems/single-number-iii/
 * 执行用时：8ms 在所有Go提交中击败了27.51%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了17.03%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNumber(nums))
	}
}
