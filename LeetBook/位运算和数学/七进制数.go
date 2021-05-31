package LeetBook

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var isMinus bool
	var ans string
	if num < 0 {
		isMinus = true
		num = -num
	}
	for num != 0 {
		ans = strconv.Itoa(num%7) + ans
		num /= 7
	}
	if isMinus {
		ans = "-" + ans
	}
	return ans
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了88.00%的用户
**/
