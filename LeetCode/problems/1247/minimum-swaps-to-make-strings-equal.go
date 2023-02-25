package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/minimum-swaps-to-make-strings-equal/
//------------------------Leetcode Problem 1247------------------------
func minimumSwap(s1 string, s2 string) int {
	xy, yx := 0, 0
	for i := range s1 {
		if s1[i] < s2[i] {
			xy++
		}
		if s1[i] > s2[i] {
			yx++
		}
	}
	if (xy+yx)%2 == 1 {
		return -1
	}
	return xy/2 + yx/2 + xy%2 + yx%2
}

//------------------------Leetcode Problem 1247------------------------
/*
 * https://leetcode.cn/problems/minimum-swaps-to-make-strings-equal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了30.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		ans := minimumSwap(strs[0], strs[1])
		fmt.Printf("%d\n", ans)
	}
}
