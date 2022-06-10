package main

/* 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
 * 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
**/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		low, high := 0, len(matrix[i])-1
		for low <= high {
			mid := low + (high-low)/2
			if target > matrix[i][mid] {
				low = mid + 1
			} else if target < matrix[i][mid] {
				high = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

func main() {
	a := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	println("target=5,", findNumberIn2DArray(a, 5))
	println("target=20,", findNumberIn2DArray(a, 20))
}

/*
 * 执行用时：32ms 在所有Go提交中击败了55.81%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了73.43%的用户
**/
