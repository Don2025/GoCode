package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func fraction(cont []int) []int {
	up := 1
	down := cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		up = cont[i]*down + up
		up, down = down, up
	}
	return []int{down, up}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cont := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", fraction(cont)[0])
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.75%的用户
**/
