package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

/*
 * 执行用时：24ms 在所有Go提交中击败了33.13%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了86.23%的用户
**/
func moveZeroes(nums []int) {
	var idx int
	for i := range nums {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for idx < len(nums) {
		nums[idx] = 0
		idx++
	}
}

/*
 * 执行用时：24ms 在所有Go提交中击败了33.13%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了7.83%的用户
func moveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		moveZeroes(nums)
		for _, x := range nums {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}
