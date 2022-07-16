package _1

type MovingAverage struct {
	Queue []int
	Size  int
	Sum   float64
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, size, 0}
}

func (ma *MovingAverage) Next(val int) float64 {
	if len(ma.Queue) == ma.Size {
		ma.Sum -= float64(ma.Queue[0])
		ma.Queue = ma.Queue[1:]
	}
	ma.Queue = append(ma.Queue, val)
	ma.Sum += float64(val)
	return ma.Sum / float64(len(ma.Queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

/*
 * 执行用时：8ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了49.66%的用户
**/
