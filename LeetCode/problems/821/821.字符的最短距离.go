package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i, j := 0, -n; i < n; i++ {
		ans[i] = n
		if s[i] == c {
			ans[i] = 0
			j = i
		} else {
			ans[i] = i - j
		}
	}
	for i, j := n-1, 2*n; i >= 0; i-- {
		if s[i] == c {
			j = i
		} else if x := j - i; x < ans[i] {
			ans[i] = x
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		s, c := arr[0], arr[1][0]
		ans := shortestToChar(s, c)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了64.20%的用户
**/
