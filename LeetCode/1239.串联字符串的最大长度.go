package main

import (
	"bufio"
	"os"
	"strings"
)

func maxLength(arr []string) int {
	var backtrace func(int, int) int
	backtrace = func(idx, bit int) int {
		var ans int
		if idx == len(arr) {
			return ans
		}
		tmp := bit
		for _, x := range arr[idx] {
			t := 1 << (x - 'a')
			if bit&t > 0 {
				return backtrace(idx+1, tmp)
			}
			bit |= t
		}
		return max(len(arr[idx])+backtrace(idx+1, bit), backtrace(idx+1, tmp))
	}
	return backtrace(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(maxLength(strings.Fields(input.Text())))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了86.11%的用户
**/
