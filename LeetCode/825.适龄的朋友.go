package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var check func(x, y int) bool
	check = func(x, y int) bool {
		if 2*y <= x+14 { // y <= 0.5*x+7
			return false
		}
		if y > x {
			return false
		}
		if y > 100 && x < 100 {
			return false
		}
		return true
	}
	var l, r, ans int
	for i := range ages {
		for l < i && !check(ages[l], ages[i]) {
			l++
		}
		if r < i {
			r = i
		}
		for r < len(ages) && check(ages[r], ages[i]) {
			r++
		}
		if r > l {
			ans += r - l - 1
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ages := stringArrayToIntArray(strings.Fields(input.Text()))
		println(numFriendRequests(ages))
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
 * 执行用时：76ms 在所有Go提交中击败了10.34%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了82.76%的用户
**/
