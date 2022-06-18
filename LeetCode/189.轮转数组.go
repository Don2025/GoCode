package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	var reverse func([]int)
	reverse = func(nums []int) {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	reverse(nums[:n-k])
	reverse(nums[n-k:])
	reverse(nums)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		rotate(nums, k)
		for _, x := range nums {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：24ms 在所有Go提交中击败了91.47%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了5.07%的用户
**/
