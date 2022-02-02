package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var ans int
	for i := 0; i < n; i++ {
		if informTime[i] != 0 {
			continue
		}
		var sum int
		for j := i; manager[j] != -1; j = manager[j] {
			sum += informTime[manager[j]]
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：436ms 在所有Go提交中击败了6.67%的用户
 * 占用内存：9.5MB 在所有Go提交中击败了91.11%的用户
**/
