package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/IDBivT/
//------------------------剑指 Offer II Problem 85------------------------
func generateParenthesis(n int) (ans []string) {
	arr := make([]byte, 0)
	var dfs func(int, int)
	dfs = func(left, right int) {
		if left+right == n<<1 {
			if left == right {
				ans = append(ans, string(arr))
			}
			return
		}
		if left < right || left > n {
			return
		}
		arr = append(arr, '(')
		dfs(left+1, right)
		arr = arr[:len(arr)-1]
		arr = append(arr, ')')
		dfs(left, right+1)
		arr = arr[:len(arr)-1]
	}
	dfs(0, 0)
	return
}

//------------------------剑指 Offer II Problem 85------------------------
/*
 * https://leetcode.cn/problems/IDBivT/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了93.89%的用户
**/
func main() {
	// You can  use Example Testcases
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a number between 1 and 8:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", generateParenthesis(n))
		Printf("Input a number:")
	}
}
