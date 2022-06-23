package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//------------------------Leetcode Problem 22------------------------
// https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(string, int, int)
	dfs = func(s string, l int, r int) {
		if l > n || r > n || r > l {
			return
		}
		if l == n && r == n {
			ans = append(ans, s)
			return
		}
		dfs(s+"(", l+1, r)
		dfs(s+")", l, r+1)
	}
	dfs("", 0, 0)
	return ans
}

//------------------------Leetcode Problem 22------------------------
/*
 * https://leetcode.cn/problems/generate-parentheses/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了34.04%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ans := generateParenthesis(n)
		fmt.Printf("%v\n", ans)
	}
}
