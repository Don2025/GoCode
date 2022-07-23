package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/occurrences-after-bigram/
//------------------------Leetcode Problem 1078------------------------
func findOcurrences(text string, first string, second string) []string {
	arr := strings.Fields(text)
	var ans []string
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] == first && arr[i+1] == second {
			ans = append(ans, arr[i+2])
		}
	}
	return ans
}

//------------------------Leetcode Problem 1078------------------------
/*
 * https://leetcode.cn/problems/occurrences-after-bigram/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了56.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		scanner.Scan()
		var first, second string
		Sscanf(scanner.Text(), "%s %s", &first, &second)
		Printf("Output: %v\n", findOcurrences(text, first, second))
	}
}
