package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/divide-a-string-into-groups-of-size-k/
//------------------------Leetcode Problem 2138------------------------
func divideString(s string, k int, fill byte) []string {
	if t := len(s) % k; t != 0 {
		for i := k - t; i > 0; i-- {
			s += string(fill)
		}
	}
	var ans []string
	for i := 0; i < len(s); i += k {
		ans = append(ans, s[i:i+k])
	}
	return ans
}

//------------------------Leetcode Problem 2138------------------------
/*
 * https://leetcode.cn/problems/divide-a-string-into-groups-of-size-k/
 * 执行用时：0ms 在所有Go提交中击败了95.17%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了69.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		fill := byte(scanner.Text()[0])
		Printf("Output: %v\n", divideString(s, k, fill))
	}
}
