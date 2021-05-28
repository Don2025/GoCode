package LeetBook

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

/*
 * 执行用时：32ms 在所有Go提交中击败了55.81%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了73.43%的用户
**/
