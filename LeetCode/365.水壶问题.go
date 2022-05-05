package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity+jug2Capacity == targetCapacity
	}
	return targetCapacity%gcd(jug1Capacity, jug2Capacity) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(canMeasureWater(arr[0], arr[1], arr[2]))
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
 * 执行用时：100ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了79.78%的用户
**/
