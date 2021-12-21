package main

import (
	"bufio"
	"os"
	"strconv"
)

func dayOfYear(date string) int {
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:])
	if m > 2 && (y%400 == 0 || (y%4 == 0 && y%100 != 0)) {
		days[2]++
	}
	var cnt int
	for i := 1; i < m; i++ {
		cnt += days[i]
	}
	cnt += d
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(dayOfYear(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：5MB 在所有Go提交中击败了95.56%的用户
**/
