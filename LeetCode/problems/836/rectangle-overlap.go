package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/rectangle-overlap/
//------------------------Leetcode Problem 836------------------------
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0] || rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
}

//------------------------Leetcode Problem 836------------------------
/*
 * https://leetcode.cn/problems/rectangle-overlap/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了33.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rec1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		rec2 := utils.StringToInts(scanner.Text())
		Printf("Output: %v", isRectangleOverlap(rec1, rec2))
	}
}
