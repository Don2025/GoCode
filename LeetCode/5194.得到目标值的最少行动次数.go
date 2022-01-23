package main

import (
	"bufio"
	"os"
	"strconv"
)

func minMoves(target int, maxDoubles int) int {
	var cnt int
	for target > 1 && maxDoubles > 0 {
		if target&1 == 0 {
			target >>= 1
			maxDoubles--
		} else {
			target--
		}
		cnt++
	}
	cnt += target - 1
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		target, _ := strconv.Atoi(input.Text())
		input.Scan()
		maxDoubles, _ := strconv.Atoi(input.Text())
		println(minMoves(target, maxDoubles))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
