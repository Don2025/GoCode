package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		adjust(nums, i)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}

func adjust(nums []int, i int) {
	for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		println(firstMissingPositive(utils.StringToInts(scanner.Text())))
	}
}

/*
 * 执行用时：112ms 在所有Go提交中击败了36.10%的用户
 * 占用内存：22MB 在所有Go提交中击败了10.59%的用户
**/
