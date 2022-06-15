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

