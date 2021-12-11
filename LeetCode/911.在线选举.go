package main

type TopVotedCandidate struct {
	Tops, Times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var tops []int
	votes := map[int]int{}
	var vs, p int
	for i := 0; i < len(times); i++ {
		votes[persons[i]]++
		if vs <= votes[persons[i]] {
			vs = votes[persons[i]]
			p = persons[i]
		}
		tops = append(tops, p)
	}
	return TopVotedCandidate{tops, times}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.Times)
	for l < r {
		m := l + (r-l)>>1
		if this.Times[m] <= t {
			l = m + 1
		} else {
			r = m
		}
	}
	return this.Tops[l-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

/*
 * 执行用时：224ms 在所有Go提交中击败了57.89%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了100.00%的用户
**/
