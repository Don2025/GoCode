package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findMinMoves(machines []int) int {
	var total int
	for _, x := range machines {
		total += x
	}
	n := len(machines)
	if total%n != 0 {
		return -1
	}
	avg := total / n
	var ans, sum int
	for _, x := range machines {
		x -= avg
		sum += x
		ans = max(ans, max(abs(sum), x))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		machines := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findMinMoves(machines))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：4ms 在所有Go提交中击败了100.0%的用户
 * 占用内存：3.99MB 在所有Go提交中击败了83.31%的用户
**/
