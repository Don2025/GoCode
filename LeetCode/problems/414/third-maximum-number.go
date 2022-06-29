package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

// https://leetcode.cn/problems/third-maximum-number/
//------------------------Leetcode Problem 414------------------------
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

//------------------------Leetcode Problem 414------------------------
/*
 * https://leetcode.cn/problems/third-maximum-number/
 * 执行用时：4ms 在所有Go提交中击败了91.62%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了53.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", thirdMax(nums))
	}
}
