package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/unique-email-addresses/
//------------------------Leetcode Problem 929------------------------
func numUniqueEmails(emails []string) int {
	emailSet := make(map[string]bool)
	for _, email := range emails {
		i := strings.IndexByte(email, '@')
		local := strings.Split(email[:i], "+")[0]  // 去掉本地名第一个加号之后的部分
		local = strings.ReplaceAll(local, ".", "") // 去掉本地名中所有的句点
		emailSet[local+email[i:]] = true
	}
	return len(emailSet)
}

//------------------------Leetcode Problem 929------------------------
/*
 * https://leetcode.cn/problems/unique-email-addresses/
 * 执行用时：8ms 在所有Go提交中击败了58.49%的用户
 * 占用内存：5.4MB 在所有Go提交中击败了69.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		emails := strings.Fields(scanner.Text())
		Printf("Output: %d\n", numUniqueEmails(emails))
	}
}
