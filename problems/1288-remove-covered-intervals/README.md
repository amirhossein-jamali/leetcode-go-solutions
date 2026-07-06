---
id: 1288
title: "Remove Covered Intervals"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/remove-covered-intervals/"
contest: "Biweekly Contest 15"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Sorting"
go_concepts:
  - "sort.Slice with custom comparator"
  - "In-place sorting of input slice"
  - "One-pass scan after sorting"
  - "Table-driven tests with slice cloning"
tags:
  - leetcode
  - go
  - array
  - sorting
  - intervals
  - biweekly-contest-15
---

# 1288. Remove Covered Intervals

## Problem Link

LeetCode: `https://leetcode.com/problems/remove-covered-intervals/`

## Difficulty

Medium

## Problem Topics

* Array
* Sorting

## What to Know Before Solving

General concepts:

* Interval coverage: `[a, b)` is covered by `[c, d)` when `c <= a` and `b <= d`
* Sorting intervals by start, then by end descending for equal starts
* One-pass maximum tracking after sorting
* Why a bucket array of size `100001` is wasteful when `n <= 1000`

Go concepts:

* `sort.Slice` with a custom comparator
* In-place sorting without copying the input slice
* Cloning `[][]int` in tests because the solution mutates the input

## Problem Description

Given an array `intervals` where `intervals[i] = [li, ri]` represents the half-open interval `[li, ri)`, remove all intervals that are covered by another interval in the list.

Return the number of remaining intervals.

## Function Signature

Expected LeetCode function signature:

```go
func removeCoveredIntervals(intervals [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
intervals = [[1,4],[3,6],[2,8]]
```

Output:

```text
2
```

Explanation:

```text
Interval [3,6] is covered by [2,8], so it is removed.
```

### Example 2

Input:

```text
intervals = [[1,4],[2,3]]
```

Output:

```text
1
```

Explanation:

```text
Interval [2,3] is covered by [1,4], so it is removed.
```

## Constraints

```text
1 <= intervals.length <= 1000
intervals[i].length == 2
0 <= li < ri <= 100000
All intervals are unique.
```

## Approach

Sort intervals by start ascending. When starts are equal, sort by end descending.

Scan once from left to right while tracking `maxEnd`, the largest right endpoint seen among non-covered intervals so far.

If the current interval's end is greater than `maxEnd`, it is not covered, so count it and update `maxEnd`. Otherwise it is covered by a previous interval.

## Algorithm

1. Sort `intervals` in place with `sort.Slice`.
2. Initialize `remaining = 0` and `maxEnd = -1`.
3. For each interval `[start, end]`:
   * If `end > maxEnd`, increment `remaining` and set `maxEnd = end`.
4. Return `remaining`.

## Why This Works

### Why sorting order matters

After sorting by start ascending, every previous interval has `previousStart <= currentStart`.

For coverage we only need to check whether some previous interval also has `previousEnd >= currentEnd`.

`maxEnd` stores the maximum right endpoint among intervals already counted as non-covered.

The end-descending tie-breaker is necessary for equal starts. For example, `[1,5]` must come before `[1,4]`; otherwise `[1,4]` may be counted incorrectly before the wider interval is seen.

### Correctness intuition

If `currentEnd <= maxEnd`, some earlier non-covered interval starts no later than the current interval and extends at least as far to the right, so the current interval is covered.

If `currentEnd > maxEnd`, no previous interval can cover the current one, because all previous intervals start no later and none extends this far right.

### Alternative bucket approach

A counting approach over coordinates `0..100000` can run in `O(n + U)` time with `U = 100000`, but it uses `O(U)` memory.

With `n <= 1000`, sort plus scan is the preferred practical solution: `O(n log n)` time and `O(1)` auxiliary space aside from sort overhead.

## Complexity Analysis

Let:

```text
n = len(intervals)
```

### Time Complexity

```text
O(n log n)
```

Sorting dominates. The scan is `O(n)`.

### Space Complexity

```text
O(1)
```

or `O(log n)` depending on Go sort internals. No auxiliary `O(n)` or `O(U)` structure is used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Sort the input slice in place; do not allocate a copy.
* Use `maxEnd = -1` so the first interval is always counted.
* Compare ends with `>` rather than `>=` because equal ends on different starts mean coverage.

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Equal-start tie-breaking
* No covered intervals
* Multiple covered intervals in one list
* Same end with different starts
* Single interval at maximum coordinate

## Edge Cases

Important cases to consider:

* Single interval
* Nested intervals with equal starts
* Touching intervals such as `[1,2]` and `[2,3]` that do not cover each other
* Many intervals covered by one wide interval
* Maximum coordinate `100000`

## Notes

* The solution mutates the input slice because LeetCode allows in-place sorting.
* Tests clone intervals before calling the function.
* Greedy local decisions without sorting do not work; coverage depends on global interval relationships.
