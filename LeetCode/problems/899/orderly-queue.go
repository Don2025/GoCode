package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/orderly-queue/
//------------------------Leetcode Problem 899------------------------
func orderlyQueue(s string, k int) string {
	if k == 1 {
		ans := s
		for i := 1; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}

//------------------------Leetcode Problem 899------------------------
/*
 * https://leetcode.cn/problems/orderly-queue/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了80.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", orderlyQueue(s, k))
	}
}
