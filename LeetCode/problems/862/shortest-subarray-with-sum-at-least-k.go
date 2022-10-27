package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
//------------------------Leetcode Problem 862------------------------
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}
	ans := n + 1
	deque := []int{}
	for i, curSum := range prefix {
		for len(deque) > 0 && curSum >= k+prefix[deque[0]] {
			ans = min(ans, i-deque[0])
			deque = deque[1:]
		}
		for len(deque) > 0 && curSum <= prefix[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if ans < n+1 {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 862------------------------
/*
 * https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
 * 执行用时：136ms 在所有Go提交中击败了65.91%的用户
 * 占用内存：9.4MB 在所有Go提交中击败了39.77%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", shortestSubarray(nums, k))
	}
}
