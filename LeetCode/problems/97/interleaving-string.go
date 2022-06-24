package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/interleaving-string/
//------------------------Leetcode Problem 97------------------------
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if (n1 + n2) != n3 {
		return false
	}
	f := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		f[i] = make([]bool, n2+1)
	}
	f[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n1][n2]
}

//------------------------Leetcode Problem 97------------------------
/*
 * https://leetcode.cn/problems/interleaving-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了79.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", isInterleave(arr[0], arr[1], arr[2]))
	}
}
