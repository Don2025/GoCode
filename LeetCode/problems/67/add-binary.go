package main

import (
	"bufio"
	. "fmt"
	"math/big"
	"os"
	"strings"
)

//https://leetcode.cn/problems/add-binary/
//------------------------Leetcode Problem 67------------------------
func addBinary(a string, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)
	ai.Add(ai, bi)
	return ai.Text(2)
}

//------------------------Leetcode Problem 67------------------------
/*
 * https://leetcode.cn/problems/add-binary/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了87.74%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", addBinary(arr[0], arr[1]))
	}
}
