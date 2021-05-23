package LeetBook

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= (num - 1)
		cnt++
	}
	return cnt
}

/*
 * 执行用时：4ms 在所有Go提交中击败了9.13%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
