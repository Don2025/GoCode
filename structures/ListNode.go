package structures

// ListNode is a singly-linked list node, which is consists of two data members:
// Val is the int value data we are keeping track of at this node;
// Next is the next ListNode in this chain.
type ListNode struct {
	Val  int
	Next *ListNode
}
