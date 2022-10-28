package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/sum-of-subarray-minimums/
//------------------------Leetcode Problem 907------------------------
func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i, x := range arr {
		for len(stack) > 0 && x <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n - i
		} else {
			right[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	for i := range arr {
		ans = (ans + left[i]*right[i]*arr[i]) % mod
	}
	return
}

//------------------------Leetcode Problem 907------------------------
/*
 * https://leetcode.cn/problems/sum-of-subarray-minimums/
 * 执行用时：64ms 在所有Go提交中击败了23.91%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了12.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", sumSubarrayMins(arr))
	}
}
