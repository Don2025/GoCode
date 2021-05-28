package LeetBook

func replaceSpace(s string) string {
	str := make([]rune, 0)
	for _, it := range s {
		if it == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, it)
		}
	}
	return string(str)
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了30.44%的用户
**/
