package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/gray-code/
//------------------------Leetcode Problem 89------------------------
// Generation process of gray code: G(i)=i^(i/2)
func grayCode(n int) []int {
	var ans []int
	for i := 0; i < (1 << n); i++ {
		ans = append(ans, i^i>>1)
	}
	return ans
}

//------------------------Leetcode Problem 89------------------------
/*
 * https://leetcode.cn/problems/gray-code/
 * 执行用时：8ms 在所有Go提交中击败了98.16%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了73.80%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", grayCode(n))
	}
}
