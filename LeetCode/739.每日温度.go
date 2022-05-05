package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				ans[i] = j - i
				break
			} else if ans[j] == 0 {
				ans[i] = 0
				break
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		tempratures := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := dailyTemperatures(tempratures)
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
 * 执行用时：620ms 在所有Go提交中击败了13.52%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了38.00%的用户
**/
