---
id: 1840
title: "Maximum Building Height"
difficulty: "Hard"
level: "Advanced"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-building-height/"
contest: "Weekly Contest 238"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Math"
  - "Sorting"
go_concepts:
  - "Custom struct slices"
  - "slices.SortFunc"
  - "Two-pass constraint propagation"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - math
  - sorting
  - weekly-contest-238
---

# 1840. Maximum Building Height

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-building-height/`

## Difficulty

Hard

## Problem Topics

* Array
* Math
* Sorting

## What to Know Before Solving

General concepts:

* Adjacent buildings can differ by at most 1 in height
* Building 1 is fixed at height 0
* Restrictions only cap heights; they do not force exact values
* The answer is the maximum peak height achievable between two known limits

Go concepts:

* Building a helper slice with custom structs
* Sorting with `slices.SortFunc`
* Forward and backward scans to tighten constraints
* Table-driven tests with copied input slices

## Problem Description

You want to build `n` buildings in a line labeled from 1 to n.

Rules:

* Each height is a non-negative integer
* Building 1 must have height 0
* Adjacent buildings differ by at most 1
* Each restriction `[idi, maxHeighti]` means building `idi` must have height less than or equal to `maxHeighti`

Building 1 never appears in `restrictions`.

Return the maximum possible height of the tallest building.

## Function Signature

Expected LeetCode function signature:

```go
func maxBuilding(n int, restrictions [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 5, restrictions = [[2,1],[4,1]]
```

Output:

```text
2
```

Explanation:

```text
Heights [0,1,2,1,2] are valid, and the tallest building has height 2.
```

### Example 2

Input:

```text
n = 6, restrictions = []
```

Output:

```text
5
```

Explanation:

```text
Heights [0,1,2,3,4,5] are valid, and the tallest building has height 5.
```

### Example 3

Input:

```text
n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
```

Output:

```text
5
```

Explanation:

```text
Heights [0,1,2,3,3,4,4,5,4,3] are valid, and the tallest building has height 5.
```

## Constraints

```text
2 <= n <= 10^9
0 <= restrictions.length <= min(n - 1, 10^5)
2 <= idi <= n
idi is unique
0 <= maxHeighti <= 10^9
```

## Approach

Both solutions use the same core idea:

1. Sort limits by building position
2. Tighten heights from left to right using the +1 per step rule
3. Tighten heights from right to left
4. Compute the maximum peak between every adjacent pair of limits

The peak between two consecutive limits at positions `p1` and `p2` with tightened heights `h1` and `h2` is:

```text
(h1 + h2 + (p2 - p1)) / 2
```

The difference is how the left and right boundaries are represented.

### Approach 1: Boundary-Based (`maxBuilding`)

This version builds a new slice of `{position, height}` limits.

It adds two explicit boundary points:

* Building 1 with height 0
* Building n with height n - 1

Then it sorts all limits, merges duplicate positions by keeping the smaller height, and runs the two tightening passes on that unified slice.

Why this version is easier to read:

* Every segment uses the same loop logic
* Building 1 and building n are treated like any other limit
* Empty restrictions are handled naturally by the two boundaries alone

Trade-offs:

* Time: `O(m log m)`
* Extra space: `O(m)`
* Does not modify the input `restrictions` slice

This is the function submitted to LeetCode.

### Approach 2: In-Place (`maxBuildingInPlace`)

This version works directly on the `restrictions` slice.

Instead of creating boundary entries, it handles the edges differently:

* Left boundary: start the forward pass from position 1 with height 0
* Right boundary: after the forward pass, initialize `maxHeight` with `lastHeight + (n - lastPosition)`
* Left peak: after the backward pass, also check `(firstHeight + firstPosition - 1) / 2`

Why this version is useful:

* No auxiliary limit slice
* Lower auxiliary memory usage
* Same asymptotic time complexity

Trade-offs:

* Time: `O(m log m)`
* Extra space: `O(1)` excluding sort internals
* Sorts and mutates the input `restrictions` slice
* Edge handling is split across three places instead of one unified loop

### Comparison

| Aspect | `maxBuilding` | `maxBuildingInPlace` |
| --- | --- | --- |
| Input mutation | No | Yes |
| Extra slice | Yes | No |
| Boundary handling | Explicit `[1, 0]` and `[n, n-1]` | Implicit formulas |
| Empty restrictions | Handled by boundaries | Early return `n - 1` |
| Restriction at building `n` | Merge with right boundary | Already present in input |
| Readability | Higher | Lower |
| LeetCode submission | Yes | No |

## Algorithm

### Boundary-Based Version

1. Build a slice of `{position, height}` limits from restrictions plus the two boundaries.
2. Sort by position and merge duplicate positions by keeping the minimum height.
3. Scan left to right and tighten each height using the +1 per step rule.
4. Scan right to left, tighten again, and track the maximum peak between each adjacent pair.
5. Return the maximum peak.

### In-Place Version

1. If `restrictions` is empty, return `n - 1`.
2. Sort `restrictions` by building position.
3. Scan left to right from building 1 with height 0 and tighten each restriction.
4. Initialize `maxHeight` using the last restriction and building `n`.
5. Scan right to left, tighten again, and track peaks between adjacent restrictions.
6. Also check the peak between building 1 and the first restriction.
7. Return the maximum peak.

## Why This Works

Any valid height sequence must satisfy both the slope rule and every restriction.

The forward pass computes the strongest upper bound each limit can have based on everything to its left. The backward pass does the same from the right. After both passes, every limit is as tight as possible without violating any rule.

Between two consecutive limits, the best peak is the midpoint of the two slopes coming from both sides. Checking every adjacent pair is enough because the optimal tallest building must lie on one of these segments.

## Complexity Analysis

Let `m` be the number of restrictions.

Both versions have the same time complexity.

### Time Complexity

```text
O(m log m)
```

Sorting dominates because both scans are linear in the number of limits.

### Space Complexity

Boundary-based version:

```text
O(m)
```

An auxiliary slice stores the boundary points plus all restrictions.

In-place version:

```text
O(1)
```

No extra limit slice is allocated. This does not count the internal memory used by sorting.

## Code

Both Go implementations are available in:

```text
solution.go
```

### LeetCode Submission: `maxBuilding`

Important implementation details:

* Boundary points make the first and last segments explicit
* Duplicate positions are merged after sorting so a restriction at building `n` does not conflict with the right boundary
* The input `restrictions` slice is not modified

```go
func maxBuilding(n int, restrictions [][]int) int {
	limits := make([]limit, 0, len(restrictions)+2)
	limits = append(limits, limit{position: 1, height: 0})

	for _, restriction := range restrictions {
		limits = append(limits, limit{
			position: restriction[0],
			height:   restriction[1],
		})
	}

	limits = append(limits, limit{position: n, height: n - 1})

	slices.SortFunc(limits, func(a, b limit) int {
		return a.position - b.position
	})

	merged := limits[:0]
	for _, current := range limits {
		if len(merged) > 0 && merged[len(merged)-1].position == current.position {
			if current.height < merged[len(merged)-1].height {
				merged[len(merged)-1].height = current.height
			}
			continue
		}
		merged = append(merged, current)
	}
	limits = merged

	for i := 1; i < len(limits); i++ {
		distance := limits[i].position - limits[i-1].position
		maxReachable := limits[i-1].height + distance

		if limits[i].height > maxReachable {
			limits[i].height = maxReachable
		}
	}

	maxHeight := 0

	for i := len(limits) - 2; i >= 0; i-- {
		distance := limits[i+1].position - limits[i].position
		maxReachable := limits[i+1].height + distance

		if limits[i].height > maxReachable {
			limits[i].height = maxReachable
		}

		peak := (limits[i].height + limits[i+1].height + distance) / 2

		if peak > maxHeight {
			maxHeight = peak
		}
	}

	return maxHeight
}
```

### Alternative: `maxBuildingInPlace`

Important implementation details:

* Works directly on `restrictions`
* Handles the right edge with `lastHeight + (n - lastPosition)`
* Handles the left edge with `(firstHeight + firstPosition - 1) / 2`
* Returns early when `restrictions` is empty

```go
func maxBuildingInPlace(n int, restrictions [][]int) int {
	if len(restrictions) == 0 {
		return n - 1
	}

	slices.SortFunc(restrictions, func(a, b []int) int {
		return a[0] - b[0]
	})

	previousPosition := 1
	previousHeight := 0

	for i := 0; i < len(restrictions); i++ {
		position := restrictions[i][0]
		maxReachable := previousHeight + position - previousPosition

		if restrictions[i][1] > maxReachable {
			restrictions[i][1] = maxReachable
		}

		previousPosition = position
		previousHeight = restrictions[i][1]
	}

	last := restrictions[len(restrictions)-1]
	maxHeight := last[1] + n - last[0]

	for i := len(restrictions) - 2; i >= 0; i-- {
		current := restrictions[i]
		next := restrictions[i+1]

		distance := next[0] - current[0]
		maxReachable := next[1] + distance

		if restrictions[i][1] > maxReachable {
			restrictions[i][1] = maxReachable
		}

		peak := (restrictions[i][1] + next[1] + distance) / 2

		if peak > maxHeight {
			maxHeight = peak
		}
	}

	firstPeak := (restrictions[0][1] + restrictions[0][0] - 1) / 2

	if firstPeak > maxHeight {
		maxHeight = firstPeak
	}

	return maxHeight
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All LeetCode examples
* Both `maxBuilding` and `maxBuildingInPlace`
* Minimum `n` without restrictions
* A restriction on the last building
* A restriction at building `n` that must merge with the right boundary
* Large `n` values
* Restrictions with very large height caps that are still limited by slope

## Edge Cases

Important cases to consider:

* Empty restrictions
* A restriction on building `n`
* A very tight restriction near building 1
* Restrictions with `maxHeighti` much larger than `n`
* Large `n` up to 10^9

## Notes

* Both versions are correct when edge handling is done carefully.
* The boundary-based version needs duplicate-position merging when a restriction exists at building `n`.
* The in-place version avoids that issue because building `n` is never duplicated, but it mutates the input.
* For LeetCode submission, prefer `maxBuilding` because it keeps the input unchanged and makes every segment follow the same logic.
* Integer overflow is not an issue in Go on 64-bit platforms because the largest intermediate sums are below 3 * 10^9.
