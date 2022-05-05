package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		letters := []byte(input.Text())
		var target byte
		fmt.Scan(&target)
		fmt.Printf("%c\n", nextGreatestLetter(letters, target))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了73.50%的用户
**/
