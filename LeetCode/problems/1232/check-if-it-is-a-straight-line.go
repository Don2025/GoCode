package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/check-if-it-is-a-straight-line/
//------------------------Leetcode Problem 1232------------------------
func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[1][1]-coordinates[0][1])*(coordinates[i][0]-coordinates[0][0]) != (coordinates[i][1]-coordinates[0][1])*(coordinates[1][0]-coordinates[0][0]) {
			return false
		}
	}
	return true
}

/*
 * https://leetcode.cn/problems/check-if-it-is-a-straight-line/
 * 执行用时：4ms 在所有Go提交中击败了95.71%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了31.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		coordinates := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			coordinates[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", checkStraightLine(coordinates))
	}
}
