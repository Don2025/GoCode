package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var cnt int
	for _, x := range costs {
		if coins >= x {
			coins -= x
		} else {
			break
		}
		cnt++
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		costs := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		coins, _ := strconv.Atoi(input.Text())
		println(maxIceCream(costs, coins))
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
 * 执行用时：232ms 在所有Go提交中击败了50.45%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了82.28%的用户
**/
