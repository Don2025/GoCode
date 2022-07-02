package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/distribute-candies/
//------------------------Leetcode Problem 575------------------------
func distributeCandies(candyType []int) int {
	mp := make(map[int]bool)
	n := len(candyType) >> 1
	for i := 0; i < len(candyType); i++ {
		mp[candyType[i]] = true
		if len(mp) >= n {
			return n
		}
	}
	return len(mp)
}

//------------------------Leetcode Problem 575------------------------
/*
 * https://leetcode.cn/problems/distribute-candies/
 * 执行用时：140ms 在所有Go提交中击败了96.71%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了75.66%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		candyType := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", distributeCandies(candyType))
	}
}
