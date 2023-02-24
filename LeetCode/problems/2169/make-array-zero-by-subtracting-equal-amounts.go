package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/
//------------------------Leetcode Problem 2357------------------------
func minimumOperations(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if num > 0 {
			m[num] = true
		}
	}
	return len(m)
}

//------------------------Leetcode Problem 2357------------------------
/*
 * https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了22.86%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		fmt.Printf("%d\n", minimumOperations(nums))
	}
}
