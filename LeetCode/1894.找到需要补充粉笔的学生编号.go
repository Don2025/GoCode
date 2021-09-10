package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func chalkReplacer(chalk []int, k int) int {
	sum := make([]int, len(chalk))
	for i := 0; i < len(chalk); i++ {
		if i == 0 {
			sum[i] = chalk[i]
		} else {
			sum[i] = sum[i-1] + chalk[i]
		}
	}
	t := k % sum[len(sum)-1]
	l, r := 0, len(sum)
	for l < r {
		m := l + (r-l)>>1
		if sum[m] <= t {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		chalk := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(chalkReplacer(chalk, k))
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
 * 执行用时：120ms 在所有Go提交中击败了94.23%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了55.77%的用户
**/
