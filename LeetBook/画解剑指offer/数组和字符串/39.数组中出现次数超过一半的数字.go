package main

//非空数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
func majorityElement(nums []int) int {
	var cnt, ans int
	for _, x := range nums {
		if cnt == 0 {
			ans = x
		}
		if ans == x {
			cnt++
		} else {
			cnt--
		}
	}
	return ans
}

func main() {
	a := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	println(majorityElement(a))
}

/*
 * 执行用时：20ms 在所有Go提交中击败了86.32%的用户
 * 占用内存：6MB 在所有Go提交中击败了98.89%的用户
**/
