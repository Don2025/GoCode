package main

import (
	"container/ring"
	"fmt"
)

type Player struct {
	position int
	alive    bool
}

func main() {
	//sum当前总人数, 数到k出圈
	var sum, k int
	_, err := fmt.Scanf("%d %d", &sum, &k)
	if err != nil {
		return
	}
	// 初始化所有玩家的信息
	r := ring.New(sum)
	for i := 1; i <= sum; i++ {
		r.Value = &Player{i, true}
		r = r.Next()
	}
	cnt, deadCnt := 1, 0 //cnt当前报数, deadCnt死亡人数
	virgin := true
	for deadCnt < sum { //留一个活的
		r = r.Next()
		//活人开始数数
		if r.Value.(*Player).alive {
			cnt++
		}
		//报到k的人出局
		if cnt == k {
			r.Value.(*Player).alive = false
			if virgin {
				r.Value.(*Player).alive = false
				fmt.Printf("%d", r.Value.(*Player).position)
				virgin = false
			} else {
				fmt.Printf(",%d", r.Value.(*Player).position)
			}
			deadCnt++
			cnt = 0
		}
	}
	println()
}
