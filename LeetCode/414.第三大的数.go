package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func thirdMax(nums []int) int {
	var cnt int
	// 注意测试样例 [1,2,-2147483648],预期结果是-2147483647,所以不能用MinInt32
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, x := range nums {
		if x == max1 || x == max2 || x == max3 {
			continue
		}
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
			cnt++
		} else if x > max2 {
			max3 = max2
			max2 = x
			cnt++
		} else if x > max3 {
			max3 = x
			cnt++
		}
	}
	if cnt < 3 {
		return max1
	}
	return max3
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(thirdMax(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：4ms 在所有Go提交中击败了91.62%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了53.89%的用户
**/
