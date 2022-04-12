package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func numberOfLines(widths []int, s string) []int {
	row, col := 1, 0
	for _, x := range s {
		col += widths[x-'a']
		if col > 100 {
			col = widths[x-'a']
			row++
		}
	}
	return []int{row, col}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		widths := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		s := input.Text()
		ans := numberOfLines(widths, s)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了68.18%的用户
**/
