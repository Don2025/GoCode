package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)
	for _, x := range arr {
		m[x]++
	}
	if m[0]&1 == 1 {
		return false
	}
	v := make([]int, 0, len(m))
	for x := range m {
		v = append(v, x)
	}
	sort.Slice(v, func(i, j int) bool { return abs(v[i]) < abs(v[j]) })
	for _, x := range v {
		if m[x<<1] < m[x] {
			return false
		}
		m[x*2] -= m[x]
	}
	return true
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
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(canReorderDoubled(arr))
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
 * 执行用时：68ms 在所有Go提交中击败了90.00%的用户
 * 占用内存：7MB 在所有Go提交中击败了83.33%的用户
**/
