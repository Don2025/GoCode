package LeetCode

func isPowerOfFour(n int) bool {
	if n < 4 {
		return n == 1
	} else if n%4 == 0 {
		return isPowerOfFour(n / 4)
	} else {
		return false
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了71.20%的用户
**/
