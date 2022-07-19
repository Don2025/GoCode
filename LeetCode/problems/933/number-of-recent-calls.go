package _33

// https://leetcode.cn/problems/number-of-recent-calls/
//------------------------Leetcode Problem 933------------------------

type RecentCounter []int

func Constructor() (r RecentCounter) {
	return
}

func (r *RecentCounter) Ping(t int) int {
	*r = append(*r, t)
	for (*r)[0] < t-3000 {
		*r = (*r)[1:]
	}
	return len(*r)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//------------------------Leetcode Problem 933------------------------
/*
 * https://leetcode.cn/problems/number-of-recent-calls/
 * 执行用时：100ms 在所有Go提交中击败了81.21%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了42.28%的用户
**/
