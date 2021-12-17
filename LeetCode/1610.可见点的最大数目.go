package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			points[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		angle, _ := strconv.Atoi(input.Text())
		input.Scan()
		location := stringArrayToIntArray(strings.Fields(input.Text()))
		println(visiblePoints(points, angle, location))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：416ms 在所有Go提交中击败了18.18%的用户
 * 占用内存：23.6MB 在所有Go提交中击败了9.09%的用户
**/
