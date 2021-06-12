package main

import "fmt"

func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		cnt += int(num & 1)
		num >>= 1
	}
	return cnt
}

func main() {
	for {
		var num uint32
		fmt.Scan(&num)
		fmt.Println(hammingWeight(num))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.84%的用户
**/
