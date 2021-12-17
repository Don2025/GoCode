package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numWaterBottles(numBottles int, numExchange int) int {
	cnt := numBottles
	for numBottles >= numExchange {
		num := numBottles / numExchange
		cnt += num
		numBottles = numBottles%numExchange + num
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		numBottles, numExchange := arr[0], arr[1]
		println(numWaterBottles(numBottles, numExchange))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.00%的用户
**/
