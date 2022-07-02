package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/find-k-closest-elements/
//------------------------Leetcode Problem 658------------------------
func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1-k
	for l <= r {
		mid := l + (r-l)>>1
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return arr[l : l+k]
}

//------------------------Leetcode Problem 658------------------------
/*
 * https://leetcode.cn/problems/find-k-closest-elements/
 * 执行用时：40ms 在所有Go提交中击败了34.01%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了73.28%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", findClosestElements(arr, k, x))
	}
}
