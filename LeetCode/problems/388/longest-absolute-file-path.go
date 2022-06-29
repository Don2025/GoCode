package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/longest-absolute-file-path/
//------------------------Leetcode Problem 388------------------------
func lengthLongestPath(input string) int {
	n := len(input)
	var ans int
	level := make([]int, n+1)
	for i := 0; i < n; {
		// 检测当前文件的深度
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		// 统计当前文件名的长度
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++ // 跳过换行符
		if depth > 1 {
			length += level[depth-1] + 1
		}
		if isFile {
			ans = max(ans, length)
		} else {
			level[depth] = length
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 388------------------------
/*
 * https://leetcode.cn/problems/longest-absolute-file-path/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了73.91%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", lengthLongestPath(scanner.Text()))
	}
}
