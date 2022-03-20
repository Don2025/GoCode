package main

import (
	"bufio"
	"os"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		t := input.Text()
		println(findTheDifference(s, t))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了84.13%的用户
**/
