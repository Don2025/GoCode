package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func jump(nums []int) int {
	var ans, end, maxDis int
	for i := 0; i < len(nums)-1; i++ {
		maxDis = max(maxDis, i+nums[i])
		if i == end {
			end = maxDis
			ans++
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Println(jump(nums))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了99.39%的用户
 * 占用内存：5.5MB 在所有Go提交中击败了100.00%的用户
**/
