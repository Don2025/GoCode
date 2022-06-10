package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, num := range nums {
		if x, ok := hashTable[target-num]; ok {
			return []int{x, i}
		}
		hashTable[num] = i
	}
	return []int{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%v\n", twoSum(nums, target))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了96.83%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了57.00%的用户
**/
