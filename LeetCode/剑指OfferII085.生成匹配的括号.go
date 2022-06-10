package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateParenthesis(n int) []string {
	var ans []string
	var backtrack func([]string, string, int, int)
	backtrack = func(ans []string, cur string, open int, close int) {
		if len(cur) == n<<1 {
			ans = append(ans, cur)
			return
		}
		if open < n {
			cur += "("
			backtrack(ans, cur, open+1, close)
			cur
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := generateParenthesis(n)
		fmt.Printf("%v\n", ans)
	}
}
