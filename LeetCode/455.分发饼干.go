package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findContentChildren(g []int, s []int) int {
	var ans, idx int
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= g[idx] {
			ans++
			idx++
			if idx >= len(g) {
				break
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		g := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		s := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findContentChildren(g, s))
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
 * 执行用时：40ms 在所有Go提交中击败了46.37%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了49.67%的用户
**/
