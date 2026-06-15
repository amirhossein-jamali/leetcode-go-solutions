---
id: 2095
title: "Delete the Middle Node of a Linked List"
difficulty: "Medium"
level: "Senior"
platform: "LeetCode"
link: "https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/"
contest: "Weekly Contest 270"
status: "Solved"
language: "Go"
topics:
  - "Senior"
  - "Linked List"
  - "Two Pointers"
go_concepts:
  - "Pointers and struct types"
  - "Singly linked lists"
  - "Slow and fast pointers"
  - "Rewiring Next pointers"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - senior
  - linked-list
  - two-pointers
  - weekly-contest-270
---

# 2095. Delete the Middle Node of a Linked List

## Problem Link

LeetCode: `https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/`

## Difficulty

Medium

## Problem Topics

* Senior
* Linked List
* Two Pointers

## What to Know Before Solving

General concepts:

* How a singly linked list connects nodes with `Next` pointers
* The middle index for length `n` is `⌊n / 2⌋` using 0-based indexing
* Using two pointers at different speeds to reach the middle node
* Deleting a node from a singly linked list by skipping it with the predecessor’s `Next`

Go concepts:

* Defining a `ListNode` struct with `Val` and `Next *ListNode`
* Keeping a `prev` pointer while advancing `slow` and `fast`
* Special case when `n = 1` (the only node is the middle; the list becomes empty)

## Problem Description

You are given the `head` of a linked list. Delete the middle node and return the head of the modified list.

The middle node for a list of size `n` is the `⌊n / 2⌋`th node from the start (0-based).

## Function Signature

Expected LeetCode function signature:

```go
func deleteMiddle(head *ListNode) *ListNode {

}
```

## Examples

### Example 1

Input:

```text
head = [1,3,4,7,1,2,6]
```

Output:

```text
[1,3,4,1,2,6]
```

Explanation:

```text
n = 7, middle index is 3, value 7 is removed.
```

### Example 2

Input:

```text
head = [1,2,3,4]
```

Output:

```text
[1,2,4]
```

Explanation:

```text
n = 4, middle index is 2, value 3 is removed.
```

### Example 3

Input:

```text
head = [2,1]
```

Output:

```text
[2]
```

Explanation:

```text
n = 2, middle index is 1, value 1 is removed.
```

## Constraints

```text
The number of nodes in the list is in the range [1, 10^5].
1 <= Node.val <= 10^5
```

## Approach

Use slow and fast pointers starting at `head`. When `fast` reaches the end of the list, `slow` lands on the middle node. Track `prev` as the node before `slow`, then set `prev.Next = slow.Next` to remove the middle.

If the list has only one node, that node is the middle; removing it yields an empty list, so return `nil`.

## Algorithm

1. If `head` is `nil` or `head.Next` is `nil`, return `nil` (length 1).
2. Initialize `slow` and `fast` to `head`, and `prev` to `nil`.
3. While `fast` and `fast.Next` are not `nil`, set `prev` to `slow`, advance `slow` by one and `fast` by two.
4. Set `prev.Next` to `slow.Next` to unlink the middle node.
5. Return `head`.

## Why This Works

The fast pointer moves twice as fast as the slow pointer. When `fast` can no longer advance by two, `slow` is exactly at the node whose index is `⌊n/2⌋` for this traversal pattern, which matches the problem’s definition of the middle. Linking `prev` to `slow.Next` removes that node without changing nodes before the middle.

For `n = 1`, the early return avoids using `prev`, which would be invalid.

## Complexity Analysis

Let `n` be the number of nodes in the list.

### Time Complexity

```text
O(n)
```

The slow/fast scan visits each node at most a constant number of times.

### Space Complexity

```text
O(1)
```

Only a constant number of pointers is used.

## Code

The Go solution is available in:

```text
solution.go
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

## Edge Cases

* One node: return `nil`
* Two nodes: delete the second node
* Odd and even lengths to confirm the correct `⌊n/2⌋` index
