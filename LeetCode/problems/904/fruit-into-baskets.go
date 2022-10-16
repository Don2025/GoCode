package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/fruit-into-baskets/
//------------------------Leetcode Problem 904------------------------
func totalFruit(fruits []int) int {
	cnt := make(map[int]int)
	var l, ans int
	for r, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[l]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 904------------------------
/*
 * https://leetcode.cn/problems/fruit-into-baskets/
 * 执行用时：112ms 在所有Go提交中击败了26.60%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了63.40%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fruits := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", totalFruit(fruits))
	}
}
