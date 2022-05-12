package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{i, j}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		numbers := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := twoSum(numbers, target)
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
 * 执行用时：4ms 在所有Go提交中击败了67.33%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了100.00%的用户
**/
