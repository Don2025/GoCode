package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/QA2IGt/
// ------------------------剑指 Offer II Problem 113------------------------
func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		ans = append(ans, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 113------------------------
/*
 * https://leetcode.cn/problems/QA2IGt/
 * 执行用时：8ms 在所有Go提交中击败了95.40%的用户
 * 占用内存：6MB 在所有Go提交中击败了60.92%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a number:")
	for scanner.Scan() {
		numCourses, _ := strconv.Atoi(scanner.Text())
		Printf("Input the number of rows of matrix:")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		prerequisites := make([][]int, n)
		for i := range prerequisites {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			prerequisites[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findOrder(numCourses, prerequisites))
		Printf("Input a number:")
	}
}
