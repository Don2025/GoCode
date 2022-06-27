package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/h-index/
//------------------------Leetcode Problem 274------------------------
func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if h <= citations[i] {
			return h
		}
	}
	return 0
}

//------------------------Leetcode Problem 274------------------------
/*
 * https://leetcode.cn/problems/h-index/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了97.32%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		citations := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", hIndex(citations))
	}
}
