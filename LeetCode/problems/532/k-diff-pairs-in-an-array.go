package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/k-diff-pairs-in-an-array/
//------------------------Leetcode Problem 532------------------------
func findPairs(nums []int, k int) int {
	visited := make(map[int]bool)
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			m[num-k] = true
		}
		if _, ok := visited[num+k]; ok {
			m[num] = true
		}
		visited[num] = true
	}
	return len(m)
}

//------------------------Leetcode Problem 532------------------------
/*
 * https://leetcode.cn/problems/k-diff-pairs-in-an-array/
 * 执行用时：8ms 在所有Go提交中击败了92.77%的用户
 * 占用内存：6MB 在所有Go提交中击败了40.96%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findPairs(nums, k))
		Printf("Input a line of numbers separated by space:")
	}
}
