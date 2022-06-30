package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/hamming-distance/
//------------------------Leetcode Problem 461------------------------
func hammingDistance(x int, y int) int {
	t, sum := x^y, 0
	for t != 0 {
		sum += t & 1
		t >>= 1
	}
	return sum
}

//------------------------Leetcode Problem 461------------------------
/*
 * https://leetcode.cn/problems/hamming-distance/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了68.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		Printf("Output: %v\n", hammingDistance(x, y))
	}
}
