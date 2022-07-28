package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/SNJvJP/
// ------------------------LeetCode Cup Problem 29------------------------
func orchestraLayout(num int, xPos int, yPos int) int {
	circle := min(xPos, num-xPos-1, yPos, num-yPos-1) //先确定这个位置在第几圈,circle从0开始计数,即第1圈为circle=0
	n := num - 1 - 2*circle                           //n记录第circle圈的边长,边长为正方形边长-1,每圈过后边长减2
	count := num*num - (n+1)*(n+1)
	if xPos == circle && yPos < num-1-circle {
		count += yPos - circle + 1
	} else if yPos == num-1-circle && xPos < num-1-circle {
		count += n + xPos - circle + 1
	} else if xPos == num-1-circle && yPos > circle {
		count += 2*n + num - circle - yPos
	} else {
		count += 3*n + num - circle - xPos
	}
	if count%9 == 0 {
		return 9
	}
	return count % 9
}

func min(a ...int) int {
	val := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < val {
			val = a[i]
		}
	}
	return val
}

// ------------------------LeetCode Cup Problem 29------------------------
/*
 * https://leetcode.cn/problems/SNJvJP/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了16.22%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var num, xPos, yPos int
		Sscanf(scanner.Text(), "%d %d %d", &num, &xPos, &yPos)
		Printf("Output: %d\n", orchestraLayout(num, xPos, yPos))
	}
}
