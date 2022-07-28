package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/stone-game-ix/
//------------------------Leetcode Problem 2029------------------------
func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)
	for _, x := range stones {
		cnt[x%3]++
	}
	if cnt[0]&1 == 0 {
		return cnt[1] >= 1 && cnt[2] >= 1
	}
	return abs(cnt[1]-cnt[2]) > 2
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 2029------------------------
/*
 * https://leetcode.cn/problems/stone-game-ix/
 * 执行用时：220ms 在所有Go提交中击败了30.77%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stones := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", stoneGameIX(stones))
	}
}
