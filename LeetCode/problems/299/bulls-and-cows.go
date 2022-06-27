package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/bulls-and-cows/
//------------------------Leetcode Problem 299------------------------
func getHint(secret string, guess string) string {
	var bulls, cows int
	var cntS, cntG [10]int
	for i := range secret {
		if guess[i] == secret[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		fmt.Printf("Output: %v\n", getHint(arr[0], arr[1]))
	}
}

//------------------------Leetcode Problem 299------------------------
/*
 * https://leetcode.cn/problems/bulls-and-cows/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了88.89%的用户
**/
