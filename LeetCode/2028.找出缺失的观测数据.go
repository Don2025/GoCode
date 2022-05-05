package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func missingRolls(rolls []int, mean int, n int) []int {
	var sum int
	for _, x := range rolls {
		sum += x
	}
	remain := mean*(n+len(rolls)) - sum
	var ans []int
	if remain > n*6 || remain < n {
		return ans
	}
	ans = make([]int, n)
	var idx int
	for remain > 0 {
		ans[idx%n]++
		remain--
		idx++
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		rolls := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		mean, _ := strconv.Atoi(input.Text())
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		ans := missingRolls(rolls, mean, n)
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
 * 执行用时：144ms 在所有Go提交中击败了46.15%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了76.92%的用户
**/
