package main

// https://leetcode.cn/problems/construct-quad-tree/
//------------------------Leetcode Problem 427------------------------

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func(int, int, int) *Node
	dfs = func(x int, y int, offset int) *Node {
		if offset <= 0 {
			return nil
		}
		for i := x; i < x+offset; i++ {
			for j := y; j < y+offset; j++ {
				if grid[i][j] != grid[x][y] {
					t := offset >> 1
					return &Node{
						false, false,
						dfs(x, y, t),
						dfs(x, y+t, t),
						dfs(x+t, y, t),
						dfs(x+t, y+t, t),
					}
				}
			}
		}
		return &Node{grid[x][y] != 0, true, nil, nil, nil, nil}
	}
	return dfs(0, 0, len(grid))
}

//------------------------Leetcode Problem 427------------------------
/*
 * https://leetcode.cn/problems/construct-quad-tree/
 * 执行用时：8ms 在所有Go提交中击败了76.19%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了76.19%的用户
**/
