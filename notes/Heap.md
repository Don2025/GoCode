```go
// https://leetcode.cn/problems/g5c51o/
//-------------------------剑指 Offer II Problem 60------------------------

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

```





```go
// https://leetcode.cn/problems/qn8gGX/
// ------------------------剑指 Offer II Problem 61------------------------
type pair struct{ i, j int }
type hp struct {
	Data         []pair
	Nums1, Nums2 []int
}

func (h hp) Len() int { return len(h.Data) }
func (h hp) Less(i, j int) bool {
	a, b := h.Data[i], h.Data[j]
	return h.Nums1[a.i]+h.Nums2[a.j] < h.Nums1[b.i]+h.Nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }
func (h *hp) Push(v interface{}) { h.Data = append(h.Data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.Data; v := a[len(a)-1]; h.Data = a[:len(a)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.Data = append(h.Data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(h, pair{i, j + 1})
		}
	}
	return
}
```





```go
// https://leetcode.cn/problems/xx4gT2/
// ------------------------剑指 Offer II Problem 76------------------------

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *intHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func findKthLargest(nums []int, k int) int {
	hp := &intHeap{}
	for _, x := range nums {
		heap.Push(hp, x)
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	return (*hp)[0]
}
```



```go
// https://leetcode.cn/problems/vvXgSW/
// ------------------------剑指 Offer II Problem 78------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type listHeap []*ListNode

func (h listHeap) Len() int            { return len(h) }
func (h listHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h listHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *listHeap) Push(v interface{}) { *h = append(*h, v.(*ListNode)) }
func (h *listHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	h := new(listHeap)
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	head := new(ListNode)
	cur := head
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		node = node.Next
		if node != nil {
			heap.Push(h, node)
		}
		cur = cur.Next
	}
	return head.Next
}
```



```go
// https://leetcode.cn/problems/qn8gGX/
// ------------------------剑指 Offer II Problem 61------------------------
type pair struct{ i, j int }
type hp struct {
	Data         []pair
	Nums1, Nums2 []int
}

func (h hp) Len() int { return len(h.Data) }
func (h hp) Less(i, j int) bool {
	a, b := h.Data[i], h.Data[j]
	return h.Nums1[a.i]+h.Nums2[a.j] < h.Nums1[b.i]+h.Nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }
func (h *hp) Push(v interface{}) { h.Data = append(h.Data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.Data; v := a[len(a)-1]; h.Data = a[:len(a)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.Data = append(h.Data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(h, pair{i, j + 1})
		}
	}
	return
}
```



```go
// https://leetcode.cn/problems/k-closest-points-to-origin/
//------------------------Leetcode Problem 973------------------------
type pair struct {
	Dist  int
	Point []int
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].Dist > h[j].Dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosest(points [][]int, k int) [][]int {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h)
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].Dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	var ans [][]int
	for _, p := range h {
		ans = append(ans, p.Point)
	}
	return ans
}
```



```go
// https://leetcode.cn/problems/maximum-number-of-eaten-apples/
//------------------------Leetcode Problem 1705------------------------

type Apple struct {
	Date int
	Num  int
}

type minHeap []Apple

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Date < h[j].Date }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(Apple)) }
func (h *minHeap) Pop() interface{} {
	ans := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ans
}

func eatenApples(apples []int, days []int) int {
	h := minHeap{}
	var cnt int
	for i := 0; i < len(days) || len(h) != 0; i++ {
		if i < len(days) && apples[i] != 0 {
			heap.Push(&h, Apple{i + days[i], apples[i]})
		}
		for len(h) != 0 && h[0].Date == i {
			heap.Pop(&h)
		}
		if len(h) != 0 && h[0].Num > 0 {
			cnt++
			h[0].Num--
			if h[0].Num == 0 {
				heap.Pop(&h)
			}
		}
	}
	return cnt
}
```



```go
// https://leetcode.cn/problems/seat-reservation-manager/
//------------------------Leetcode Problem 1845------------------------
type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Top() int            { return h[0] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type SeatManager struct {
	Seats *hp
}

func Constructor(n int) SeatManager {
	m := SeatManager{&hp{}}
	heap.Init(m.Seats)
	for i := 1; i <= n; i++ {
		heap.Push(m.Seats, i)
	}
	return m
}

func (this *SeatManager) Reserve() int {
	x := this.Seats.Top()
	heap.Pop(this.Seats)
	return x
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.Seats, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
```



```go
// https://leetcode.cn/problems/p0NxJO/
// ------------------------LeetCode Cup Problem 30------------------------
type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func magicTower(nums []int) int {
	sum, att, ans := 1, 0, 0 // sum当前血量, att被换到最后房间的怪物伤害, ans结果
	h := new(hp)
	for _, x := range nums {
		if x < 0 {
			heap.Push(h, x)
		}
		sum += x
		for sum <= 0 {
			if len(*h) != 0 {
				ans++
				t := heap.Pop(h).(int)
				sum -= t
				att -= t
			} else {
				return -1
			}
		}
	}
	if sum > att {
		return ans
	}
	return -1
}
```



```go
// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
// ------------------------剑指 Offer I Problem 40------------------------

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	h := make(IntHeap, k)
	hp := &h
	copy(*hp, arr[:k+1])
	heap.Init(hp)
	for i := k; i < len(arr); i++ {
		if arr[i] < h[0] {
			heap.Pop(hp)
			heap.Push(hp, arr[i])
		}
	}
	return h
}
```

