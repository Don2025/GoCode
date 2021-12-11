package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1]+1
		}
	}
	right := make([]int, n)
	for i := n-2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			right[i] = right[i+1]+1
		}
	}
	var ans []int
	for i := 0; i < n; i++ {
		if left[i]>=time && right[i]>=time {
			ans = append(ans, i)
		}
	}
	return ans
}

/* Time Limit Exceeded
func goodDaysToRobBank(security []int, time int) []int {
	var ans []int
	for i := time; i < len(security)-time; i++ {
		flag := true
		for j := i-time; j < i; j++ {
			if security[j] < security[j+1] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		for j := i; j < i+time; j++ {
			if security[j] > security[j+1] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		security := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		time, _ := strconv.Atoi(input.Text())
		ans := goodDaysToRobBank(security, time)
		for _, x := range ans {
			fmt.Printf("%d ", x )
		}
		fmt.Println()
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
 * 执行用时：108ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：10MB 在所有Go提交中击败了100.00%的用户
**/
