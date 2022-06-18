package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if visited[num] {
			return true
		}
		visited[num] = true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		println(containsDuplicate(nums))
	}
}

/*
 * 执行用时：92ms 在所有Go提交中击败了17.05%的用户
 * 占用内存：8.8MB 在所有Go提交中击败了15.83%的用户
**/
