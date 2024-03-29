package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/poor-pigs/
//------------------------Leetcode Problem 458------------------------
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	var cnt int
	maxRound := minutesToTest/minutesToDie + 1
	for int(math.Pow(float64(maxRound), float64(cnt))) < buckets {
		cnt++
	}
	return cnt
}

//------------------------Leetcode Problem 458------------------------
/*
 * https://leetcode.cn/problems/poor-pigs/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了72.73%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		buckets, _ := strconv.Atoi(arr[0])
		minutesToDie, _ := strconv.Atoi(arr[1])
		minutesToTest, _ := strconv.Atoi(arr[2])
		Printf("Output: %v\n", poorPigs(buckets, minutesToDie, minutesToTest))
	}
}
