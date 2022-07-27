package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/kth-missing-positive-number/
//------------------------Leetcode Problem 1539------------------------
func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return k + l
}

//------------------------Leetcode Problem 1539------------------------
/*
 * https://leetcode.cn/problems/kth-missing-positive-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了13.79%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findKthPositive(arr, k))
	}
}
