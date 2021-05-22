package LeetCode

import "sort"

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

/*
 * 执行用时：40ms 在所有Go提交中击败了46.37%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了49.67%的用户
**/
