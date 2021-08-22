package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func escapeGhosts(ghosts [][]int, target []int) bool {
	dis := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		if abs(g[0]-target[0])+abs(g[1]-target[1]) <= dis {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ghosts := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			ghosts[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		target := stringArrayToIntArray(strings.Fields(input.Text()))
		println(escapeGhosts(ghosts, target))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：0MB 在所有Go提交中击败了100.00%的用户
**/
