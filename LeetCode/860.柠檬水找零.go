package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, x := range bills {
		if x == 5 {
			five++
		} else if x == 10 {
			five--
			ten++
		} else if ten > 0 {
			five--
			ten--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bills := stringArrayToIntArray(strings.Fields(input.Text()))
		println(lemonadeChange(bills))
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
 * 执行用时：92ms 在所有Go提交中击败了87.81%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了73.38%的用户
**/
