package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/day-of-the-week/
//------------------------Leetcode Problem 1185------------------------
func dayOfTheWeek(day int, month int, year int) string {
	if month == 1 || month == 2 {
		month += 12
		year--
	}
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return week[(day+2*month+3*(month+1)/5+year+year/4-year/100+year/400)%7]
}

//------------------------Leetcode Problem 1185------------------------
/*
 * https://leetcode.cn/problems/day-of-the-week/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var day, month, year int
		Sscanf(scanner.Text(), "%d %d %d", &day, &month, &year)
		Printf("Output: %v\n", dayOfTheWeek(day, month, year))
	}
}
