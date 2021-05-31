package LeetCode

var dic = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := dic[s[i]]
		if i < n-1 && value < dic[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

/*
 * 执行用时：12ms 在所有Go提交中击败了41.95%的用户
 * 占用内存：3MB 在所有Go提交中击败了100.00%的用户
**/
