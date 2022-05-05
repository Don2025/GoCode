package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > 0 && ty > 0 {
		if sx == tx && sy == ty {
			return true
		}
		if tx > ty {
			tx -= ty * max((tx-sx)/ty, 1)
		} else {
			ty -= tx * max((ty-sy)/tx, 1)
		}
	}
	return false
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
		arr := strings.Fields(input.Text())
		sx, _ := strconv.Atoi(arr[0])
		sy, _ := strconv.Atoi(arr[1])
		tx, _ := strconv.Atoi(arr[2])
		ty, _ := strconv.Atoi(arr[3])
		println(reachingPoints(sx, sy, tx, ty))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
