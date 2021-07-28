package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func hasGroupsSizeX(deck []int) bool {
	mp := make(map[int]int)
	for _, x := range deck {
		mp[x]++
	}
	var cnt int
	for _, x := range mp {
		if x > 0 {
			cnt = gcd(cnt, x)
			if cnt == 1 {
				return false
			}
		}
	}
	return cnt >= 2
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
		println(hasGroupsSizeX(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：12ms 在所有Go提交中击败了92.68%的用户
 * 占用内存：5.6MB 在所有Go提交中击败了65.85%的用户
**/
