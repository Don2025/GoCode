package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"

	"strconv"
)

// https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk/
//------------------------Leetcode Problem 1894------------------------
func chalkReplacer(chalk []int, k int) int {
	sum := make([]int, len(chalk))
	for i := 0; i < len(chalk); i++ {
		if i == 0 {
			sum[i] = chalk[i]
		} else {
			sum[i] = sum[i-1] + chalk[i]
		}
	}
	t := k % sum[len(sum)-1]
	l, r := 0, len(sum)
	for l < r {
		m := l + (r-l)>>1
		if sum[m] <= t {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

//------------------------Leetcode Problem 1894------------------------
/*
 * https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk/
 * 执行用时：120ms 在所有Go提交中击败了94.23%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了55.77%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		chalk := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", chalkReplacer(chalk, k))
	}
}
