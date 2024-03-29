package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/contains-duplicate-ii/
//------------------------Leetcode Problem 219------------------------
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i, num := range nums {
		if x, ok := dict[num]; ok && i-x <= k {
			return true
		}
		dict[num] = i
	}
	return false
}

//------------------------Leetcode Problem 219------------------------
/*
 * https://leetcode.cn/problems/contains-duplicate-ii/
 * 执行用时：104ms 在所有Go提交中击败了73.19%的用户
 * 占用内存：12MB 在所有Go提交中击败了20.06%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", containsNearbyDuplicate(nums, k))
	}
}
