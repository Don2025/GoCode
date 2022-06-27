package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/next-greater-element-ii/
//------------------------Leetcode Problem 503------------------------
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < n<<1-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}

//------------------------Leetcode Problem 503------------------------
/*
 * https://leetcode.cn/problems/next-greater-element-ii/
 * 执行用时：24ms 在所有Go提交中击败了69.44%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了84.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", nextGreaterElements(nums))
	}
}
