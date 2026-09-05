---
id: 3904
title: "Smallest Stable Index II"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/smallest-stable-index-ii/"
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

# 3904. Smallest Stable Index II

## Problem Link

LeetCode: `https://leetcode.com/problems/smallest-stable-index-ii/`

## Difficulty

Medium

## Problem Topics

* Array
* Prefix Sum

## What to Know Before Solving

General concepts:

* Instability score at index `i` is `max(nums[0..i]) - min(nums[i..n-1])`
* A stable index satisfies score `<= k`
* Finding the leftmost stable index
* How a later small suffix value can invalidate every earlier candidate
* Why `n <= 10^5` still allows a linear scan, but scanning `n` with extra arrays is unnecessary

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

`max(nums[0..i])` is the largest value among the elements from index `0` to index `i`.

`min(nums[i..n - 1])` is the smallest value among the elements from index `i` to index `n - 1`.

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

Explanation:

```text
Every index has score 3 - 1 = 2, which is greater than k = 1.
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
1 <= nums.length <= 10^5
0 <= nums[i] <= 10^9
0 <= k <= 10^9
```

## Approach

Index `i` is stable if and only if every value in `nums[i..n-1]` is at least `max(nums[0..i]) - k`.

A suffix-minimum array would make this check direct, but it uses `O(n)` extra memory. The same answer can be found in one forward pass with three integers:

* `ans`: the earliest index that may still be stable
* `mx`: prefix maximum of `nums[0..i]`
* `pm`: prefix maximum at `ans`

If a later value `nums[i]` satisfies `pm - nums[i] > k`, then the suffix starting at `ans` contains a value that is too small. Move `ans` to `i + 1`.

## Algorithm

1. Initialize `ans = 0`, `mx = nums[0]`, and `pm = nums[0]`.
2. For each index `i` from `1` to `n - 1`:
   * Update `mx` with `nums[i]`.
   * If `i == ans`, set `pm = mx`.
   * If `pm - nums[i] > k`, set `ans = i + 1`.
3. If `ans == n`, return `-1`.
4. Otherwise return `ans`.

## Why This Works

### Stability is a suffix check against a fixed prefix maximum

Index `i` is stable if and only if:

```text
max(nums[0..i]) - min(nums[i..n-1]) <= k
```

which is the same as:

```text
for every j >= i, nums[j] >= max(nums[0..i]) - k
```

So a single later value that is smaller than `pm - k` is enough to reject the current candidate.

### A violating value also rejects every skipped index

Suppose `pm - nums[i] > k` and the current candidate is `ans`. For every index `t` with `ans <= t <= i`:

* `max(nums[0..t]) >= pm`, because prefix maxima are non-decreasing
* `min(nums[t..n-1]) <= nums[i]`, because `nums[i]` still belongs to that suffix

Therefore:

```text
max(nums[0..t]) - min(nums[t..n-1]) >= pm - nums[i] > k
```

Every index from `ans` through `i` is unstable. Jumping to `i + 1` does not skip a valid answer.

### The surviving candidate is the leftmost stable index

`ans` starts at `0` and moves only when a later value proves it unstable. If the scan ends with `ans < n`, no later value invalidated it, so `ans` is stable. Because every smaller index was rejected, it is also the smallest stable index.

If `ans == n`, every index was invalidated and the answer is `-1`.

## Complexity Analysis

Let:

```text
n = len(nums)
```

### Time Complexity

```text
O(n)
```

One forward scan. This is asymptotically optimal because any later value can invalidate the current candidate, so the whole array must be read.

A suffix-minimum array is also `O(n)` time, but it is not needed.

### Space Complexity

```text
O(1)
```

Only `ans`, `mx`, and `pm` are stored. This stays constant even though `n` can be `10^5`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `pm` is refreshed only when the scan reaches the current `ans`
* `pm - nums[i] > k` is equivalent to `nums[i] + k < pm`
* Returning `-1` when `ans == n` means every index was invalidated
* The loop starts at index `1` because a one-element array is always stable for `k >= 0`

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
* A late small value that invalidates every earlier candidate
* Large values relative to `k`
* Duplicate values throughout the array

## Notes

* This is the same statement as Smallest Stable Index I, with `n` increased from `100` to `10^5`.
* A two-array prefix/suffix approach is simpler but uses `O(n)` extra memory.
* The one-pass candidate scan keeps auxiliary space constant and still runs in linear time.
