package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	for i := 0; i < 1<<n; i++ {
		var arr []int
		for j := range nums {
			if i>>j&1 > 0 {
				arr = append(arr, nums[j])
			}
		}
		ans = append(ans, append([]int(nil), arr...))
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		fmt.Printf("%v\n", subsets(nums))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了25.72%的用户
**/
