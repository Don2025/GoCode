package main

type RecentCounter []int

func Constructor() (_ RecentCounter) {
	return
}

func (rc *RecentCounter) Ping(t int) int {
	*rc = append(*rc, t)
	for (*rc)[0] < t-3000 {
		*rc = (*rc)[1:]
	}
	return len(*rc)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

/*
 * 执行用时：100ms 在所有Go提交中击败了76.43%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了27.14%的用户
**/
