package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func asteroidCollision(asteroids []int) []int {
	var ans []int
	for _, a := range asteroids {
		for a < 0 && len(ans) > 0 && ans[len(ans)-1] > 0 {
			sum := ans[len(ans)-1] + a
			if sum >= 0 {
				a = 0
			}
			if sum <= 0 {
				ans = ans[:len(ans)-1]
			}
		}
		if a != 0 {
			ans = append(ans, a)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		asteroids := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := asteroidCollision(asteroids)
		fmt.Printf("%v\n", ans)
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
 * 执行用时：12ms 在所有Go提交中击败了20.16%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了51.94%的用户
**/
