---
id: 3471
title: "Find the Largest Almost Missing Integer"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/find-the-largest-almost-missing-integer/"
contest: "Weekly Contest 439"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
go_concepts:
  - "Fixed-size array [51]int"
  - "Frequency counting"
  - "Case analysis on window size"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - hash-table
  - counting
  - weekly-contest-439
---

# 3471. Find the Largest Almost Missing Integer

## Problem Link

LeetCode: `https://leetcode.com/problems/find-the-largest-almost-missing-integer/`

## Difficulty

Easy

## Problem Topics

* Array
* Hash Table

## What to Know Before Solving

General concepts:

* A subarray is a contiguous segment of the array
* An integer `x` is almost missing if it appears in exactly one subarray of length `k`
* Values lie in `[0, 50]`, so frequencies fit in a fixed array
* The number of length-`k` windows that contain an index depends on that index's position

Go concepts:

* Counting with a stack-allocated `[51]int`
* Ranging over a fixed array, where the index is the value itself
* Splitting the solution into cases on `k`
* Writing table-driven tests with the `testing` package

## Problem Description

You are given an integer array `nums` and an integer `k`.

An integer `x` is almost missing if it appears in exactly one contiguous subarray of length `k`.

Return the largest almost missing integer. If none exists, return `-1`.

## Function Signature

Expected LeetCode function signature:

```go
func largestInteger(nums []int, k int) int {

}
```

## Examples

### Example 1

Input:

```text
nums = [3,9,2,1,7], k = 3
```

Output:

```text
7
```

`3` and `7` each appear in exactly one window of length `3`. The larger value is `7`.

### Example 2

Input:

```text
nums = [3,9,7,2,1,7], k = 4
```

Output:

```text
3
```

Only `3` appears in exactly one window of length `4`.

### Example 3

Input:

```text
nums = [0,0], k = 1
```

Output:

```text
-1
```

`0` appears in two windows of length `1`.

## Constraints

```text
1 <= nums.length <= 50
0 <= nums[i] <= 50
1 <= k <= nums.length
```

## Approach

There are `n - k + 1` windows of length `k`. How many of those windows contain a given index depends on `k`:

* If `k == n`, there is only one window: the whole array.
* If `k == 1`, each index is its own window.
* If `1 < k < n`, every interior index belongs to at least two windows. Only the first and last elements can appear in exactly one window.

Because values are in `[0, 50]`, frequencies are stored in a fixed `[51]int`.

## Algorithm

1. If `k == n`, return the maximum value in `nums`. Duplicates are still valid because they all sit in the same single window.
2. Count frequencies in `freq[0..50]`.
3. If `k == 1`, return the largest value whose frequency is exactly `1`, or `-1` if none exists.
4. Otherwise only `nums[0]` and `nums[n-1]` are candidates. Keep a candidate if it occurs exactly once in the whole array, then return the larger valid candidate, or `-1`.

## Why This Works

### `k == n`

There is one subarray of length `k`. Every distinct value in `nums` appears in that unique window, even if it repeats inside it. The largest array value is the answer.

### `k == 1`

Each position is a window. The number of windows containing `x` equals the number of times `x` occurs. The answer is the largest value with frequency `1`.

Ranging over `freq` visits keys `0..50` in order, so the last key with count `1` is the maximum unique value.

### `1 < k < n`

Index `0` belongs only to the first window. Index `n-1` belongs only to the last window. Every interior index `i` belongs to more than one window.

If `nums[0]` (or `nums[n-1]`) also appears somewhere else, it appears in at least two windows. So a candidate is valid only when its global frequency is `1`.

## Complexity Analysis

Let:

```text
n = len(nums)
```

Values are bounded by `50`, so the frequency array has constant size.

### Time Complexity

```text
O(n)
```

The array is scanned a constant number of times. The scan of `freq` is `O(1)`.

### Space Complexity

```text
O(1)
```

Only a fixed `[51]int` is used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `k == n` is handled before frequency counting
* For `k == 1`, ranging over `freq` yields the maximum unique value without an extra comparison
* For the middle case, `-1` is a valid initial answer because all values are non-negative

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* A single-element array
* `k == n` with a duplicated maximum
* `k == 1` with a unique maximum
* Both endpoints duplicated when `1 < k < n`
* Unique endpoints where the first value is larger

## Edge Cases

Important cases to consider:

* `n = 1`
* `k = 1` with no unique values
* `k = n` with duplicates
* First and last values equal
* First valid, last invalid, and the reverse

## Notes

* Sliding over every window is unnecessary once the three cases are separated.
* Interior values cannot be almost missing when `1 < k < n`.
* When `k == n`, repeating a value does not disqualify it.
