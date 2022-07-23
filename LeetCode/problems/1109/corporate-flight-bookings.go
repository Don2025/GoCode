package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/corporate-flight-bookings/
//------------------------Leetcode Problem 1109------------------------
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			ans[i-1] += booking[2]
		}
	}
	return ans
}

//------------------------Leetcode Problem 1109------------------------
/*
 * https://leetcode.cn/problems/corporate-flight-bookings/
 * 执行用时：820ms 在所有Go提交中击败了18.66%的用户
 * 占用内存：9MB 在所有Go提交中击败了70.15%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		bookings := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			bookings[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", corpFlightBookings(bookings, n))
	}
}
