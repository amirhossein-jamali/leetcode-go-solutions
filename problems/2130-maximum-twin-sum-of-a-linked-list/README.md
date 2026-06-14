---
id: 2130
title: "Maximum Twin Sum of a Linked List"
difficulty: "Medium"
level: "Senior"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/"
contest: "Biweekly Contest 69"
status: "Solved"
language: "Go"
topics:
  - "Senior"
  - "Linked List"
  - "Two Pointers"
  - "Stack"
go_concepts:
  - "Pointers and struct types"
  - "Singly linked lists"
  - "Slow and fast pointers"
  - "In-place list reversal"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - senior
  - linked-list
  - two-pointers
  - stack
  - biweekly-contest-69
---

# 2130. Maximum Twin Sum of a Linked List

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/`

## Difficulty

Medium

## Problem Topics

* Senior
* Linked List
* Two Pointers
* Stack
* Biweekly Contest 69

## What to Know Before Solving

General concepts:

* How a singly linked list connects nodes with `Next` pointers
* What it means for two nodes to be twins in an even-length list
* Using two pointers at different speeds to find the middle region
* Reversing a linked list by rewiring `Next` pointers in place

Go concepts:

* Defining a `ListNode` struct with `Val` and `Next *ListNode`
* Iterating with `for curr != nil` and saving `next` before rewiring
* Returning the new head after reversal (`prev` pattern)
* Building test lists from slices in `solution_test.go`

## Problem Description

In a linked list of size `n`, where `n` is even, the `i`th node (0-indexed) is the twin of the `(n - 1 - i)`th node for `0 <= i <= (n / 2) - 1`.

The twin sum is the sum of the values of a node and its twin.

Given the `head` of a linked list with even length, return the maximum twin sum over all such pairs.

## Function Signature

Expected LeetCode function signature:

```go
func pairSum(head *ListNode) int {

}
```

## Examples

### Example 1

Input:

```text
head = [5,4,2,1]
```

Output:

```text
6
```

Explanation:

```text
Node 0 (5) is the twin of node 3 (1): sum = 6.
Node 1 (4) is the twin of node 2 (2): sum = 6.
The maximum twin sum is 6.
```

### Example 2

Input:

```text
head = [4,2,2,3]
```

Output:

```text
7
```

Explanation:

```text
Twin sums are 4 + 3 = 7 and 2 + 2 = 4.
The maximum is 7.
```

### Example 3

Input:

```text
head = [1,100000]
```

Output:

```text
100001
```

Explanation:

```text
Only one twin pair exists: 1 + 100000 = 100001.
```

## Constraints

```text
The number of nodes in the list is an even integer in the range [2, 10^5].
1 <= Node.val <= 10^5
```

## Approach

Twin pairs are symmetric around the center of the list. The first half pairs with the second half in reverse order.

Use a slow and fast pointer so that when the fast pointer reaches the end, the slow pointer is at the start of the second half. Reverse the second half so its traversal order matches the twin pairing from the head of the first half.

Walk both halves in parallel, compute each twin sum, and track the maximum.

## Algorithm

1. Initialize `slow` and `fast` to `head`.
2. Advance `slow` one step and `fast` two steps until `fast` reaches the end. Then `slow` is at the first node of the second half.
3. Reverse the list starting at `slow` to get `second`, the head of the reversed second half.
4. Set `first` to `head` and `maxSum` to `0`.
5. While `second` is not `nil`, add `first.Val` and `second.Val`, update `maxSum` if needed, then move both pointers forward.
6. Return `maxSum`.

## Why This Works

By definition, twin indices sum to `n - 1`. Walking from the head of the original list and from the tail of the original list visits exactly those pairs. Reversing the second half makes a forward walk from the split point equivalent to walking from the original tail.

Every twin pair is visited once in the final loop, so the maximum over those sums is correct.

## Complexity Analysis

Let `n` be the number of nodes in the list.

### Time Complexity

```text
O(n)
```

Finding the middle is `O(n)`, reversing the second half is `O(n/2)`, and the final scan is `O(n/2)`.

### Space Complexity

```text
O(1)
```

Only a constant number of pointers is used besides the input list. The reversal is in place.

## Code

The Go solution is available in:

```text
solution.go
```

A common implementation uses slow/fast pointers plus iterative reversal:

```go
package maximumtwinsum

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverseList(slow)
	first := head
	maxSum := 0

	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}

		first = first.Next
		second = second.Next
	}

	return maxSum
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

## Edge Cases

Important cases to consider:

* Minimum length: two nodes
* All twin sums equal
* Maximum values from constraints (`10^5`) on both ends of a pair
* Longer lists where the maximum comes from the middle pairing
