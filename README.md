The library is under reconstruction, that is not only made the code more testable, but also improved the code quality in general and promoted code reuse in the refactored code sections.

As you can see, refactored code such as [**remove-nth-node-from-end-of-list.go**](https://github.com/Don2025/GoCode/blob/master/LeetCode/problems/19/remove-nth-node-from-end-of-list.go)：

- You can`cd ./LeetCode/problems/19` and then use `go test` to testify program.
- Function `Test19()` called to use Leetcode official example testcases, which eliminates manual input of test cases in local debugging.

- If output is not expected, which means the test case fails, this line `t.Errorf(err)` will executed.
- You also can use some customize test case from the standard input. ( I assume that all inputs are legal, when I wrote code.)
- These code is needs to be submitted to Leetcode, which between `//------------------------Leetcode Problem 19------------------------` the two notes.

```go
package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type ListNode = structures.ListNode

//------------------------Leetcode Problem 19------------------------
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
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		head := utils.StringToLinkList(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: ")
		utils.PrintLinkedList(removeNthFromEnd(head, n))
		Printf("Input a line of numbers separated by space:")
	}
}
```

Test code such as [**19_test.go**](https://github.com/Don2025/GoCode/blob/master/LeetCode/problems/19/19_test.go):

```go
package main

import (
	"errors"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"testing"
)

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
	Println("------------------------Leetcode Problem 19------------------------")
	for i, exam := range examples {
		input, expected := exam.argv19, exam.ans19
		Printf("Example %d:\n", i+1)
		Printf("Input: head = %v, n = %d\n", input.one, input.two)
		output := utils.ListNodeToInts(removeNthFromEnd(utils.IntsToListNode(input.one), input.two))
		Printf("Output: %v\n", output)
		Printf("Expected: %v\n", expected)
		if Sprintf("%v", output) == Sprintf("%v", expected) {
			Printf("There's not much of a difference.\n")
			Println("-------------------------------------------------------------------")
		} else {
			err := Sprintf("Example %d went wrong!\n", i+1)
			t.Errorf(err) // Fail
			// t.Fatalf(err) // FailNow
			// panic(errors.New(err))
		}
	}
	Println("You can input some customize test case right now.")
}
```

You can`cd ./LeetCode/problems/19` and then use `go test` to testify program [**remove-nth-node-from-end-of-list.go**](https://github.com/Don2025/GoCode/blob/master/LeetCode/problems/19/remove-nth-node-from-end-of-list.go) in Windows system:

```bash
D:\Code\GoCode\LeetCode\problems\19>go test
------------------------Leetcode Problem 19------------------------
Example 1:
Input: head = [1 2 3 4 5], n = 2
Output: [1 2 3 5]
Expected: [1 2 3 5]
There's not much of a difference.
-------------------------------------------------------------------
Example 2:
Input: head = [1], n = 1
Output: []
Expected: []
There's not much of a difference.
-------------------------------------------------------------------
Example 3:
Input: head = [1 2], n = 1
Output: [1]
Expected: [1]
There's not much of a difference.
-------------------------------------------------------------------
You can input some customize test case right now.
PASS
ok      github.com/Don2025/GoCode/LeetCode/problems/19  0.226s
```

