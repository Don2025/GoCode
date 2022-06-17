package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/duplicate-zeros/
//------------------------Leetcode Problem 1089------------------------
func duplicateZeros(arr []int) {
	var cnt int
	for _, x := range arr {
		if x == 0 {
			cnt++
		}
	}
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 0 {
			cnt--
		}
		if i+cnt < n {
			arr[i+cnt] = arr[i]
			if arr[i] == 0 && i+cnt+1 < n {
				arr[i+cnt+1] = 0
			}
		}
	}
}

//------------------------Leetcode Problem 1089------------------------
/*
 * https://leetcode.cn/problems/duplicate-zeros/
 * 执行用时：4ms 在所有Go提交中击败了92.59%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了100.00%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		duplicateZeros(nums)
		Printf("Output: %v\n", nums)
		Printf("Input a line of numbers separated by space:")
	}
}
