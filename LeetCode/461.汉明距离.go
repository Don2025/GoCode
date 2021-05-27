package LeetCode

func hammingDistance(x int, y int) int {
	t, sum := x^y, 0
	for t != 0 {
		sum += t & 1
		t >>= 1
	}
	return sum
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了68.49%的用户
**/
