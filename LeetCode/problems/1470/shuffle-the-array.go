package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shuffle-the-array/
//------------------------Leetcode Problem 1470------------------------
func shuffle(nums []int, n int) []int {
	ans := make([]int, n<<1)
	for i := 0; i < n; i++ {
		ans[i<<1] = nums[i]
		ans[i<<1+1] = nums[i+n]
	}
	return ans
}

//------------------------Leetcode Problem 1470------------------------
/*
 * https://leetcode.cn/problems/shuffle-the-array/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了66.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", shuffle(nums, n))
	}
}
