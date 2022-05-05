package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var ans []int
	endSet := make(map[int]bool)
	for _, e := range edges {
		endSet[e[1]] = true
	}
	for i := 0; i < n; i++ {
		if !endSet[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			input.Scan()
			edges[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := findSmallestSetOfVertices(n, edges)
		fmt.Printf("%v\n", ans)
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
 * 执行用时：176ms 在所有Go提交中击败了37.33%的用户
 * 占用内存：18.4MB 在所有Go提交中击败了21.33%的用户
**/
