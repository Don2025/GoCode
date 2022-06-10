package LeetBook

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, ((a & b) << 1))
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了8.22%的用户
**/
