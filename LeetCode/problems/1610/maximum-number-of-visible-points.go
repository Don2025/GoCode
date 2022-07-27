package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/maximum-number-of-visible-points/
//------------------------Leetcode Problem 1610------------------------
func visiblePoints(points [][]int, angle int, location []int) int {
	var sameCnt int
	polarDegrees := []float64{}
	for _, point := range points {
		if point[0] == location[0] && point[1] == location[1] {
			sameCnt++
			continue
		}
		degree := math.Atan2(float64(point[1]-location[1]), float64(point[0]-location[0]))
		polarDegrees = append(polarDegrees, degree)
	}
	sort.Float64s(polarDegrees)
	for _, x := range polarDegrees {
		polarDegrees = append(polarDegrees, x+2*math.Pi)
	}
	var maxCnt int
	degree := float64(angle) * math.Pi / 180
	var binarySearch func([]float64, float64) int
	binarySearch = func(nums []float64, target float64) int {
		l, r := 0, len(nums)-1
		var ans int
		for l <= r {
			mid := l + (r-l)>>1
			if target < nums[mid] {
				r = mid - 1
				ans = mid
			} else {
				l = mid + 1
			}
		}
		return ans
	}
	for i := range polarDegrees {
		maxCnt = max(maxCnt, binarySearch(polarDegrees, polarDegrees[i]+degree)-i)
	}
	return maxCnt + sameCnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1610------------------------
/*
 * https://leetcode.cn/problems/maximum-number-of-visible-points/
 * 执行用时：416ms 在所有Go提交中击败了18.18%的用户
 * 占用内存：23.6MB 在所有Go提交中击败了9.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			points[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		angle, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		location := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", visiblePoints(points, angle, location))
	}
}
