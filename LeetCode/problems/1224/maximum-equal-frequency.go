package main

func maxEqualFreq(nums []int) int {
	cnt, freq := map[int]int{}, map[int]int{}
	var maxFreq, ans int
	for i, num := range nums {
		if cnt[num] > 0 {
			freq[cnt[num]]--
		}
		cnt[num]++
		maxFreq = max(maxFreq, cnt[num])
		freq[cnt[num]]++
		if maxFreq == 1 || freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 || freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1224------------------------
/*
 * https://leetcode.cn/problems/maximum-equal-frequency/
 * 执行用时：80ms 在所有Go提交中击败了60.00%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了80.00%的用户
**/
func main() {

}
