package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/bulb-switcher/
//------------------------Leetcode Problem 319------------------------
func bulbSwitch(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}

//------------------------Leetcode Problem 319------------------------
/*
 * https://leetcode.cn/problems/bulb-switcher/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", bulbSwitch(n))
	}
}
