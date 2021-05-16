package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find(target int, array [][]int) bool {
	for i := 0; i < len(array); i++ {
		low, high := 0, len(array[i])-1
		for low <= high {
			mid := low + (high-low)/2
			if target > array[i][mid] {
				low = mid + 1
			} else if target < array[i][mid] {
				high = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

/*
 * 运行时间：7ms 超过14.98%用Go提交的代码
 * 占用内存：3160KB 超过62.17%用Go提交的代码
**/
