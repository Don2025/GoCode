package main

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	var ans []int
	if len(input) < k {
		return ans
	}
	sort.Ints(input)
	if len(input) == k {
		return input
	}
	for i := 0; i < k; i++ {
		ans = append(ans, input[i])
	}
	return ans
}

/*
 * 运行时间：4ms 超过18.44%用Go提交的代码
 * 占用内存：1308KB 超过10.24%用Go提交的代码
**/
