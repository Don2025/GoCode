package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/sum-of-digits-of-string-after-convert/
//------------------------Leetcode Problem 1945------------------------
func getLucky(s string, k int) int {
	t := []byte{}
	for _, c := range s {
		t = append(t, strconv.Itoa(int(c-'a'+1))...)
	}
	digits := string(t)
	for i := 1; i <= k && len(digits) > 1; i++ {
		sum := 0
		for _, c := range digits {
			sum += int(c - '0')
		}
		digits = strconv.Itoa(sum)
	}
	ans, _ := strconv.Atoi(digits)
	return ans
}

//------------------------Leetcode Problem 1945------------------------
/*
 * https://leetcode.cn/problems/sum-of-digits-of-string-after-convert/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了30.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n,", getLucky(s, k))
	}
}
