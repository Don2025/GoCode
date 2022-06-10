The library is under reconstruction, that is not only made the code more testable, but also improved the code quality in general and promoted code reuse in the refactored code sections.

As you can see, refactored code such as [**19.remove-nth-node-from-end-of-list.go**](https://github.com/Don2025/GoCode/blob/master/LeetCode/problems/19.remove-nth-node-from-end-of-list.go)：

- Function `Test19()` called to use Leetcode official example testcases, which eliminates manual input of test cases in local debugging.
- If output is not expected, which means the test case fails, this line `panic(errors.New(err))` will executed, so the program panic directly.
- You also can use some customize test case from the standard input.
- These code is needs to be submitted to Leetcode, which between `//------------------------Leetcode Problem 19------------------------` the two notes.

```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
	"testing"
)

type ListNode = structures.ListNode
// ------------------------Leetcode Problem 19------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
//------------------------Leetcode Problem 19------------------------
/*
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了61.93%的用户
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 * Runtime: 3 ms, faster than 48.26% of Go online submissions for Remove Nth Node From End of List.
 * Memory Usage: 2.2 MB, less than 32.91% of Go online submissions for Remove Nth Node From End of List.
**/

func main() {
	var t *testing.T = &testing.T{}
	Test19(t) // Use Example Testcases
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToLinkList(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		utils.PrintLinkedList(removeNthFromEnd(head, n))
	}
}

type ques19 struct {
	argv19
	ans19 []int
}

type argv19 struct {
	one []int
	two int
}

func Test19(t *testing.T) {
	examples := []ques19{
		{
			argv19{[]int{1, 2, 3, 4, 5}, 2},
			[]int{1, 2, 3, 5},
		},
		{
			argv19{[]int{1}, 1},
			[]int{},
		},
		{
			argv19{[]int{1, 2}, 1},
			[]int{1},
		},
	}
	fmt.Println("------------------------Leetcode Problem 19------------------------")
	for i, exam := range examples {
		input, expected := exam.argv19, exam.ans19
		fmt.Printf("Example %d:\n", i+1)
		fmt.Printf("Input: head = %v, n = %d\n", input.one, input.two)
		output := utils.LinkListToInts(removeNthFromEnd(utils.IntsToLinkList(input.one), input.two))
		fmt.Printf("Output: %v\n", output)
		fmt.Printf("Expected: %v\n", expected)
		if fmt.Sprintf("%v", output) == fmt.Sprintf("%v", expected) {
			fmt.Printf("There's not much of a difference.\n")
			fmt.Println("-------------------------------------------------------------------")
		} else {
			err := fmt.Sprintf("Example %d went wrong!\n", i+1)
			panic(errors.New(err))
		}
	}
	fmt.Println("You can input some customize test case right now.")
}
```