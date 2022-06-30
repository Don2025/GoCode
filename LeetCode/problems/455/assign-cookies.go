package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/assign-cookies/
//------------------------Leetcode Problem 455------------------------
func findContentChildren(g []int, s []int) int {
	var ans, idx int
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= g[idx] {
			ans++
			idx++
			if idx >= len(g) {
				break
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 455------------------------
/*
 * https://leetcode.cn/problems/assign-cookies/
 * 执行用时：40ms 在所有Go提交中击败了46.37%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了49.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g := utils.StringToInts(scanner.Text())
		scanner.Scan()
		s := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findContentChildren(g, s))
	}
}
