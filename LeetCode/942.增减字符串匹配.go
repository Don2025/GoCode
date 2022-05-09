package main

import (
	"bufio"
	"fmt"
	"os"
)

func diStringMatch(s string) []int {
	i, d := 0, len(s)
	var ans []int
	for _, x := range s {
		if x == 'I' {
			ans = append(ans, i)
			i++
		} else if x == 'D' {
			ans = append(ans, d)
			d--
		}
	}
	ans = append(ans, d)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := diStringMatch(input.Text())
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了80.88%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了42.65%的用户
**/
