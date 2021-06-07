## Container/ring

Container中包含了Go常用的容器类型，本篇记录`container/ring`的读后感。

### package ring

ring是环形链表中的一个元素，环没有起点和终点，指向任何ring元素的指针都是对整个环的引用。空环表示为nil环指针，而零环是一个具有零值的单元素环。

```go
package ring

type Ring struct {
  next, prev *Ring
  Value interface{}
}

// init 初始化一个单元素环
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// Next 返回非空环元素r的下一个环元素
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 返回非空环元素r的前一个环元素
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move 能将非空环r中的n%r.Len()个元素向后移动(n<0),或向前移动(n>=0),并返回该环元素
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// New 创建一个包含n个元素的环形链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// Link 能将非空环r与环元素s连接起来，使r.Next()变成环s，并返回r.Next()的原始值
/* 如果环r和环s指向同一个环，将它们连接起来就会把环元素r和环元素s之间的元素从环中移除，
 * 被删除的元素形成一个子环，返回结果是对该子环的引用。
 * 如果没有被删除的元素，返回结果仍是r.Next()的原始值，而不是nil。
**/
/* 如果环r和环s指向不同的环，则在环元素r后连接环元素s从而创建一个单独的环，
 * 返回结果是插入前的环元素r.Next()。
**/
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
    //不能在一行多重赋值，这样会打乱执行顺序
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

/* Unlink 能从非空环r中移除从r.next()开始的n%r.Len()个环元素，若n%r.Len()==0,环r不变
 * 返回值是被移除元素构成的子环。
**/
func (r *Ring) Unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len 返回环形链表r中的元素个数，执行时间与元素数成正比
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

// Do 对环r中的每个环元素依次执行函数，如果f改变环指针*r，则Do的行为是未定义的
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
```

### 用法示例

```go
package main

import (
    "container/ring"
    "fmt"
)

func main() {
    const rLen = 5
    // 创建一个长度为rLen的环形链表
    r := ring.New(rLen)
    fmt.Printf("r.Len() is %d.\n", r.Len())
    // 用几个整数初始化环元素的值
    for i := 0; i < rLen; i++ {
        r.Value = i
        r = r.Next()
    }
    // 从前往后遍历环形链表并打印环元素的值
    for i := 0; i < rLen; i++ {
        fmt.Printf("%d ", r.Value)
        r = r.Next()
    }
    // 0 1 2 3 4
    println()
    // 从后往前遍历环形链表并打印环元素的值
    for i := 0; i < rLen; i++ {
        r = r.Prev()
        fmt.Printf("%d ", r.Value)
    }
    // 4 3 2 1 0
    println()
    //使用Do调用匿名函数来打印环形链表的值
    r.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 0 1 2 3 4
    println()
    //将环指针向后移动2步
    r = r.Move(2)
    r.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 2 3 4 0 1
    println()
    //将环指针r后的第4个元素乘以5
    r.Move(4).Value = r.Move(4).Value.(int)*5
    r.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 2 3 4 0 5
    println()
    //删除环指针r后的3个元素
    res := r.Unlink(3)
    //可以看到原环形链表从r.Next()开始的3个元素被删除了
    r.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 2 5
    println()
    //Unlink的返回值是被删除的元素组成的子环
    res.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 3 4 0
    println()
    //将环形链表res插入到环形链表r的后面,返回值是插入前的r.Next()
    ans := r.Link(res)
    ans.Do(func(x interface{}) {
        fmt.Printf("%d ", x)
    })
    // 5 2 3 4 0
    println()
}
```

### 例题求解

> #### 模拟约瑟夫环
>
> 引言：据说著名犹太历史学家 Josephus有过以下的故事：在罗马人占领乔塔帕特后，39 个犹太人与Josephus及他的朋友躲到一个洞中，39个犹太人决定宁愿死也不要被敌人抓到，于是决定了一个自杀方式，41个人排成一个圆圈，由第1个人开始报数，每报数到第3人该人就必须自杀，然后再由下一个重新报数，直到所有人都自杀身亡为止。然而Josephus 和他的朋友并不想遵从。首先从一个人开始，越过k-2个人（因为第一个人已经被越过），并杀掉第k个人。接着，再越过k-1个人，并杀掉第k个人。这个过程沿着圆圈一直进行，直到最终只剩下一个人留下，这个人就可以继续活着。一开始要站在什么地方才能避免被处决？Josephus要他的朋友先假装遵从，他将朋友与自己安排在第16个与第31个位置，于是逃过了这场死亡游戏。
>
> **题目描述：n 个人围成一圈，从第一个人开始顺序编号为1到n。从第1个人从1开始报数，数到k(k>1)的人出圈。再由下一个人从1开始报数，数到k的人出圈，如此循环数下去，直到最后一个人出圈。请打印出圈的人的编号顺序，编号用英文逗号分割。**
>
> **输入样例1：41 3**
>
> **输出样例1：3,6,9,12,15,18,21,24,27,30,33,36,39,1,5,10,14,19,23,28,32,37,41,7,13,20,26,34,40,8,17,29,38,11,25,2,22,4,35,16,31**
>
> **输入样例2：8 3**
>
> **输出样例2：3,6,1,5,2,8,4,7** 

##### 环形链表解法：

```go
package main

import (
    "container/ring"
    "fmt"
)

type Player struct {
    position int
    alive bool
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
    for deadCnt < sum {  //留一个活的
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
```

##### 递归解法：

```go
package main

import "fmt"

//sum当前总人数, 数到k出圈, 最后存活人数n
func josephus(sum, k, n int ) int {
    if n == 1 {
        return (sum+k-1)%sum
    } else {
        return (josephus(sum-1, k, n-1)+k)%sum
    }
}

func main() {
    var sum, k int
    _, err := fmt.Scanf("%d %d", &sum, &k)
    if err != nil {
        return
    }
    virgin := true
    for i := 1; i <= sum; i++ {
        if virgin {
            fmt.Printf("%d", josephus(sum, k, i)+1)
            virgin = false
        } else {
            fmt.Printf(",%d", josephus(sum, k, i)+1)
        }
    }
    println()
}
```

