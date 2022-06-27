package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/single-number/
//------------------------Leetcode Problem 136------------------------
func singleNumber(nums []int) int {
	var ans int
	for _, x := range nums {
		ans ^= x
	}
	return ans
}

//------------------------Leetcode Problem 136------------------------
/*
 * https://leetcode.cn/problems/single-number/
 * 执行用时：12ms 在所有Go提交中击败了93.66%的用户
 * 占用内存：6MB 在所有Go提交中击败了78.66%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNumber(nums))
	}
}
