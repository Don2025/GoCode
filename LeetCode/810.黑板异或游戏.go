package LeetCode

func xorGame(nums []int) bool {
	var flag, n = 0, len(nums)
	for _, x := range nums {
		flag ^= x
	}
	if flag == 0 || n%2 == 0 {
		return true
	}
	return false
}

/*
 * 执行用时：8ms 在所有Go提交中击败了66.67%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了100.00%的用户
**/
