package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func openLock(deadends []string, target string) int {
	var cnt int //旋转次数
	deadendsMap, visited := make(map[string]bool), make(map[string]bool)
	for _, x := range deadends {
		deadendsMap[x] = true //记录所有“死亡点”
	}
	queue := []string{"0000"}
	for len(queue) != 0 {
		size := len(queue)          //BFS当前level的结点个数
		for i := 0; i < size; i++ { //遍历当前层的结点
			node := queue[0]    //队首结点
			queue = queue[1:]   //出队
			if node == target { //若队首结点正好是目标结点
				return cnt
			}
			if _, has := visited[node]; has { //若访问过改结点则略过
				continue
			}
			if _, has := deadendsMap[node]; has { //若遇到“死亡点”也略过
				continue
			}
			visited[node] = true //标记访问过的结点
			for j, x := range node {
				num := int(x - '0')
				//往上拧,往下拧所得的新数入队
				queue = append(queue, node[:j]+strconv.Itoa((num+1)%10)+node[j+1:]) //往上拧所得的新数入队
				queue = append(queue, node[:j]+strconv.Itoa((num+9)%10)+node[j+1:]) //往下拧所得的新数入队
			}
		}
		cnt++
	}
	return -1 //找不到目标结点则返回1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		deadends := strings.Fields(input.Text())
		input.Scan()
		target := input.Text()
		println(openLock(deadends, target))
	}
}

/*
 * 执行用时：136ms 在所有Go提交中击败了29.60%的用户
 * 占用内存：8.5MB 在所有Go提交中击败了5.05%的用户
**/
