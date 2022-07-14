package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
//------------------------Leetcode Problem 744------------------------
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if target < letters[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return letters[l%n]
}

//------------------------Leetcode Problem 744------------------------
/*
 * https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了73.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		letters := []byte(scanner.Text())
		scanner.Scan()
		var target byte
		Sscan(scanner.Text(), &target)
		Printf("%c\n", nextGreatestLetter(letters, target))
	}
}
