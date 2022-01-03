package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func dayOfTheWeek(day int, month int, year int) string {
	if month == 1 || month == 2 {
		month += 12
		year--
	}
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return week[(day+2*month+3*(month+1)/5+year+year/4-year/100+year/400)%7]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		day, month, year := arr[0], arr[1], arr[2]
		println(dayOfTheWeek(day, month, year))
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
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
