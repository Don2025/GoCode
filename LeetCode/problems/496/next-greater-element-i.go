package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/next-greater-element-i/
//------------------------Leetcode Problem 496------------------------
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		x := nums2[i]
		for len(stack) > 0 && x >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[x] = stack[len(stack)-1]
		} else {
			m[x] = -1
		}
		stack = append(stack, x)
	}
	var ans []int
	for _, x := range nums1 {
		ans = append(ans, m[x])
	}
	return ans
}

//------------------------Leetcode Problem 496------------------------
/*
 * https://leetcode.cn/problems/next-greater-element-i/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了25.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		nums2 := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", nextGreaterElement(nums1, nums2))
	}
}
