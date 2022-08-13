package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to/
//------------------------Leetcode Problem 1282------------------------
func groupThePeople(groupSizes []int) [][]int {
	groups := map[int][]int{}
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}
	var ans [][]int
	for size, group := range groups {
		for i := 0; i < len(group); i += size {
			ans = append(ans, group[i:i+size])
		}
	}
	return ans
}

//------------------------Leetcode Problem 1282------------------------
/*
 * https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to/
 * 执行用时：8ms 在所有Go提交中击败了40.91%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了59.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		groupSizes := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", groupThePeople(groupSizes))
	}
}
