package LeetCode

func checkSubarraySum(nums []int, k int) bool {
	mp := map[int]int{0: -1}
	var mod int
	for i, x := range nums {
		mod = (mod + x) % k
		if preidx, has := mp[mod]; has {
			if i-preidx >= 2 {
				return true
			}
		} else {
			mp[mod] = i
		}
	}
	return false
}

/*
 * 执行用时：228ms 在所有Go提交中击败了5.04%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了38.13%的用户
**/
