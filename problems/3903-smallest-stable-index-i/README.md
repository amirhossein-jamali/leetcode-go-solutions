---
id: 3903
title: "Smallest Stable Index I"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/smallest-stable-index-i/"
contest: "Weekly Contest 498"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Prefix Sum"
go_concepts:
  - "Single-pass candidate scan"
  - "Prefix maximum tracking"
  - "Table-driven tests with prefix-suffix reference"
tags:
  - leetcode
  - go
  - array
  - prefix-sum
  - weekly-contest-498
---

# 3903. Smallest Stable Index I

## Problem Link

LeetCode: `https://leetcode.com/problems/smallest-stable-index-i/`

## Difficulty

Easy

## Problem Topics

* Array
* Prefix Sum

## What to Know Before Solving

General concepts:

* Instability score at index `i` is `max(nums[0..i]) - min(nums[i..n-1])`
* A stable index satisfies score `<= k`
* Finding the leftmost stable index
* How a small suffix value can invalidate an earlier candidate

Go concepts:

* Tracking a moving candidate index in one forward pass
* Maintaining a running prefix maximum
* Comparing against a brute-force prefix-max / suffix-min reference in tests

## Problem Description

You are given an integer array `nums` of length `n` and an integer `k`.

For each index `i`, define its instability score as:

```text
max(nums[0..i]) - min(nums[i..n - 1])
```

An index `i` is stable if its instability score is less than or equal to `k`.

Return the smallest stable index. If none exists, return `-1`.

## Function Signature

Expected LeetCode function signature:

```go
func firstStableIndex(nums []int, k int) int {

}
```

## Examples

### Example 1

Input:

```text
nums = [5,0,1,4], k = 3
```

Output:

```text
3
```

Explanation:

```text
Index 0: 5 - 0 = 5
Index 1: 5 - 0 = 5
Index 2: 5 - 1 = 4
Index 3: 5 - 4 = 1

Index 3 is the first index with score <= 3.
```

### Example 2

Input:

```text
nums = [3,2,1], k = 1
```

Output:

```text
-1
```

### Example 3

Input:

```text
nums = [0], k = 0
```

Output:

```text
0
```

## Constraints

```text
1 <= nums.length <= 100
0 <= nums[i] <= 10^9
0 <= k <= 10^9
```

## Approach

A naive solution builds prefix maxima and suffix minima, then scans indices from left to right. That works, but this solution finds the leftmost stable index in one pass.

Maintain:

* `candidate`: the earliest index that may still be stable
* `prefixMax`: maximum of `nums[0..i]`
* `candidateMax`: prefix maximum at the current candidate

When a later value `x` satisfies `x + k < candidateMax`, the suffix starting at `candidate` contains a value small enough that instability exceeds `k`. Move `candidate` to `i + 1`.

## Algorithm

1. Initialize `candidate = 0`, `prefixMax = nums[0]`, and `candidateMax = prefixMax`.
2. For each index `i` from `1` to `n - 1`:
   * Update `prefixMax` with `nums[i]`.
   * If `i == candidate`, refresh `candidateMax`.
   * If `nums[i] + k < candidateMax`, set `candidate = i + 1`.
3. If `candidate == n`, return `-1`.
4. Otherwise return `candidate`.

## Why This Works

The instability of index `i` is large exactly when some suffix value from `i` onward is too small compared with the prefix maximum at `i`.

If any later value `x` is smaller than `candidateMax - k`, then every index still at or before that position fails. Jumping `candidate` past that small value is safe because earlier indices are already known to be unstable.

When the scan ends and `candidate < n`, no later value invalidated it, so `candidate` is the leftmost stable index.

## Complexity Analysis

Let:

```text
n = len(nums)
```

### Time Complexity

```text
O(n)
```

One forward scan.

### Space Complexity

```text
O(1)
```

Only a constant number of variables are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `candidateMax` is refreshed only when the scan reaches the current candidate
* The condition `x + k < candidateMax` is equivalent to `candidateMax - x > k`
* Returning `-1` when `candidate == n` means every index was invalidated

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* First index already stable
* Only the last index stable
* All-equal array with `k = 0`
* Comparison against a prefix-max / suffix-min reference

## Edge Cases

Important cases to consider:

* Single-element array
* No stable index
* First index stable
* Large values relative to `k`
* Duplicate values throughout the array

## Notes

* A two-array prefix/suffix approach is simpler but uses `O(n)` extra memory.
* The one-pass candidate scan keeps auxiliary space constant.
* Part II of this problem family may require larger constraints or different query patterns.
