package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/find-missing-observations/
//------------------------Leetcode Problem 2028------------------------
func missingRolls(rolls []int, mean int, n int) []int {
	var sum int
	for _, x := range rolls {
		sum += x
	}
	remain := mean*(n+len(rolls)) - sum
	var ans []int
	if remain > n*6 || remain < n {
		return ans
	}
	ans = make([]int, n)
	var idx int
	for remain > 0 {
		ans[idx%n]++
		remain--
		idx++
	}
	return ans
}

//------------------------Leetcode Problem 2028------------------------
/*
 * https://leetcode.cn/problems/find-missing-observations/
 * 执行用时：144ms 在所有Go提交中击败了46.15%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了76.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rolls := utils.StringToInts(scanner.Text())
		scanner.Scan()
		mean, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", missingRolls(rolls, mean, n))
	}
}
