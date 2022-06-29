package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/find-the-difference/
//------------------------Leetcode Problem 389------------------------
func findTheDifference(s string, t string) byte {
	var sum int
	for _, x := range s {
		sum -= int(x)
	}
	for _, x := range t {
		sum += int(x)
	}
	return byte(sum)
}

//------------------------Leetcode Problem 389------------------------
/*
 * https://leetcode.cn/problems/find-the-difference/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了84.13%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		t := scanner.Text()
		Printf("Output: %v\n", findTheDifference(s, t))
	}
}
