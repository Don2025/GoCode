package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := generateParenthesis(n)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了34.04%的用户
**/
