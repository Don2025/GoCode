package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[1][1]-coordinates[0][1])*(coordinates[i][0]-coordinates[0][0]) != (coordinates[i][1]-coordinates[0][1])*(coordinates[1][0]-coordinates[0][0]) {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		coordinates := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			coordinates[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(checkStraightLine(coordinates))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：4ms 在所有Go提交中击败了95.71%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了31.43%的用户
**/
