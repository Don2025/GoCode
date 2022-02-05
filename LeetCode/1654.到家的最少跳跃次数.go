package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func minimumJumps(forbidden []int, a int, b int, x int) int {
	type queue struct {
		pos  int
		back bool
	}
	path := make(map[int]bool)
	var m int
	for _, x := range forbidden {
		m = max(m, x)
		path[x] = true
	}
	m += a + b + x
	path[0] = true
	q := []queue{{0, false}}
	var ans int
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			t := q[i]
			if t.pos == x {
				return ans
			}
			if t.pos+a <= m && !path[t.pos+a] {
				q = append(q, queue{t.pos + a, false})
				path[t.pos+a] = true
			}
			if !t.back && t.pos-b > 0 && !path[t.pos-b] {
				q = append(q, queue{t.pos - b, true})
			}
		}
		ans++
		q = q[size:]
	}
	return -1
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
		forbidden := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(minimumJumps(forbidden, arr[0], arr[1], arr[2]))
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
 * 执行用时：8ms 在所有Go提交中击败了84.09%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了61.36%的用户
**/
