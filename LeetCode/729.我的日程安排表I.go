package main

type Seg struct {
	Start, End int
	Head, Tail *Seg
}

type MyCalendar struct {
	Root *Seg
}

func Constructor() MyCalendar {
	return MyCalendar{&Seg{-1, 0, nil, nil}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var insert func(*Seg, *Seg) bool
	insert = func(cur, fresh *Seg) bool {
		if fresh.Start < cur.End && fresh.End > cur.Start {
			return false
		}
		if fresh.End <= cur.Start {
			if cur.Head == nil {
				cur.Head = fresh
				return true
			}
			return insert(cur.Head, fresh)
		}
		if cur.Tail == nil {
			cur.Tail = fresh
			return true
		}
		return insert(cur.Tail, fresh)
	}
	return insert(this.Root, &Seg{start, end, nil, nil})
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

/*
 * 执行用时：64ms 在所有Go提交中击败了93.75%的用户
 * 占用内存：7MB 在所有Go提交中击败了86.25%的用户
**/
