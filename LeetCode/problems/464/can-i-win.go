package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/can-i-win/
//------------------------Leetcode Problem 464------------------------
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	dp := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(state int) bool {
		if x, ok := dp[state]; ok {
			return x
		}
		var sum int
		for i := 0; i < maxChoosableInteger; i++ {
			if (state >> i & 1) == 1 {
				sum += i + 1
			}
		}
		for i := 0; i < maxChoosableInteger; i++ {
			if (state >> i & 1) == 0 {
				if sum+i+1 >= desiredTotal || !dfs(state|1<<i) {
					dp[state] = true
					return true
				}
			}
		}
		dp[state] = false
		return false
	}
	return dfs(0)
}

//------------------------Leetcode Problem 464------------------------
/*
 * https://leetcode.cn/problems/can-i-win/
 * 执行用时：380ms 在所有Go提交中击败了14.29%的用户
 * 占用内存：31.4MB 在所有Go提交中击败了14.29%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		maxChoosableInteger, _ := strconv.Atoi(arr[0])
		desiredTotal, _ := strconv.Atoi(arr[1])
		Printf("%v\n", canIWin(maxChoosableInteger, desiredTotal))
	}
}
