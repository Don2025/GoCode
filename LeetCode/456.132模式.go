package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64
	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(find132pattern(nums))
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
 * 执行用时：36ms 在所有Go提交中击败了98.82%的用户
 * 占用内存：11.3MB 在所有Go提交中击败了21.18%的用户
**/
