package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			ans[i-1] += booking[2]
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		bookings := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			bookings[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		n, _ = strconv.Atoi(input.Text())
		ans := corpFlightBookings(bookings, n)
		for _, x := range ans {
			fmt.Printf("%d ", x)
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
 * 执行用时：820ms 在所有Go提交中击败了18.66%的用户
 * 占用内存：9MB 在所有Go提交中击败了70.15%的用户
**/
