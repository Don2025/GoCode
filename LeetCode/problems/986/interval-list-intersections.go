package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/interval-list-intersections/
//------------------------Leetcode Problem 986------------------------
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		l := max(firstList[i][0], secondList[j][0])
		r := min(firstList[i][1], secondList[j][1])
		if l <= r {
			ans = append(ans, []int{l, r})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 986------------------------
/*
 * https://leetcode.cn/problems/interval-list-intersections/
 * 执行用时：16ms 在所有Go提交中击败了90.71%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了30.77%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		arr1 := make([][]int, n)
		for i := range arr1 {
			scanner.Scan()
			arr1[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		arr2 := make([][]int, n)
		for i := range arr2 {
			scanner.Scan()
			arr2[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", intervalIntersection(arr1, arr2))
	}
}
