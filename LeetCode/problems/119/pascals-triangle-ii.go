package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/pascals-triangle-ii/
//------------------------Leetcode Problem 119------------------------
func getRow(rowIndex int) []int {
	var ans []int
	num := 1
	for i := 0; i <= rowIndex; i++ {
		ans = append(ans, num)
		num = num * (rowIndex - i) / (i + 1)
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rowIndex, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", getRow(rowIndex))
	}
}

//------------------------Leetcode Problem 119------------------------
/*
 * https://leetcode.cn/problems/pascals-triangle-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了64.20%的用户
**/
