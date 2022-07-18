package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/contain-virus/
//------------------------Leetcode Problem 749------------------------
func containVirus(isInfected [][]int) (ans int) {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(isInfected), len(isInfected[0])
	for {
		neighbors := []map[pair]bool{}
		var firewalls []int
		for i, row := range isInfected {
			for j, infected := range row {
				if infected != 1 {
					continue
				}
				q := []pair{{i, j}}
				neighbor := map[pair]bool{}
				firewall, idx := 0, len(neighbors)+1
				row[j] = -idx
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; x >= 0 && x < m && y >= 0 && y < n {
							if isInfected[x][y] == 1 {
								q = append(q, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								firewall++
								neighbor[pair{x, y}] = true
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}
		if len(neighbors) == 0 {
			break
		}
		var idx int
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}
		ans += firewalls[idx]
		for _, row := range isInfected {
			for j, infected := range row {
				if infected < 0 {
					if infected != -idx-1 {
						row[j] = 1
					} else {
						row[j] = 2
					}
				}
			}
		}
		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}
		if len(neighbors) == 1 {
			break
		}
	}
	return
}

//------------------------Leetcode Problem 749------------------------
/*
 * https://leetcode.cn/problems/contain-virus/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了60.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		isInfected := make([][]int, n)
		for i := range isInfected {
			scanner.Scan()
			isInfected[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %d\n", containVirus(isInfected))
	}
}
