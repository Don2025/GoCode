package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func subArrayRanges(nums []int) int64 {
	const INF int = 1e9 + 1
	var ans int
	for i := 0; i < len(nums); i++ {
		maxVal, minVal := -INF, INF
		for j := i; j < len(nums); j++ {
			maxVal = max(maxVal, nums[j])
			minVal = min(minVal, nums[j])
			ans += maxVal - minVal
		}
	}
	return int64(ans)
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

/* Time Limit Exceeded
func subArrayRanges(nums []int) int64 {
	var sum int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			sum += getRange(nums[i:j])
		}
	}
	return int64(sum)
}

//获取子数组的范围
func getRange(nums []int) int {
	maxVal, minVal := nums[0], nums[0]
	for _, x := range nums {
		maxVal = max(maxVal, x)
		minVal = min(minVal, x)
	}
	return maxVal - minVal
}
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(subArrayRanges(nums))
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
 * 执行用时：12ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了100.00%的用户
**/
