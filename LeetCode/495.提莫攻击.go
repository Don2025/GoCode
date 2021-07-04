package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := duration
	for i := len(timeSeries) - 2; i >= 0; i-- {
		if timeSeries[i]+duration < timeSeries[i+1] {
			sum += duration
		} else {
			sum += timeSeries[i+1] - timeSeries[i]
		}
	}
	return sum
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		timeSeries := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		duration, _ := strconv.Atoi(input.Text())
		println(findPoisonedDuration(timeSeries, duration))
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
 * 执行用时：44ms 在所有Go提交中击败了62.31%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了93.08%的用户
**/
