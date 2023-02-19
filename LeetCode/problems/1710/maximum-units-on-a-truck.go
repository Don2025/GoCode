package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/maximum-units-on-a-truck/
//------------------------Leetcode Problem 1710------------------------
func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, box := range boxTypes {
		if box[0] >= truckSize {
			ans += truckSize * box[1]
			break
		}
		truckSize -= box[0]
		ans += box[0] * box[1]
	}
	return
}

//------------------------Leetcode Problem 1710------------------------
/*
 * https://leetcode-cn.com/problems/maximum-units-on-a-truck/
 * 执行用时：20ms 在所有Go提交中击败了84.09%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了97.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		boxTypes := make([][]int, n)
		for i := range boxTypes {
			scanner.Scan()
			boxTypes[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		truckSize, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", maximumUnits(boxTypes, truckSize))
	}
}
