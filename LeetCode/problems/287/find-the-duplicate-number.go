package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/find-the-duplicate-number/
//------------------------Leetcode Problem 287------------------------
func findDuplicate(nums []int) int {
	var slow, fast int
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	for slow = 0; slow != fast; {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

//------------------------Leetcode Problem 287------------------------
/*
 * https://leetcode.cn/problems/find-the-duplicate-number/
 * 执行用时：76ms 在所有Go提交中击败了92.51%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了49.62%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findDuplicate(nums))
	}
}
