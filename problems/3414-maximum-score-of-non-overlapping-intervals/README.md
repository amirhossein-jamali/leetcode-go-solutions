---
id: 3414
title: "Maximum Score of Non-overlapping Intervals"
difficulty: "Hard"
level: "Senior Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/"
contest: "Weekly Contest 431"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Binary Search"
  - "Dynamic Programming"
  - "Sorting"
go_concepts:
  - "Weighted interval scheduling DP"
  - "Binary search on sorted end times"
  - "Bit-packed path reconstruction in uint64"
  - "slices.Sort on encoded rows"
  - "Rolling DP layers with slice reuse"
  - "reflect.DeepEqual in table-driven tests"
tags:
  - leetcode
  - go
  - array
  - binary-search
  - dynamic-programming
  - sorting
  - weekly-contest-431
---

# 3414. Maximum Score of Non-overlapping Intervals

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/`

## Difficulty

Hard

## Problem Topics

* Array
* Binary Search
* Dynamic Programming
* Sorting

## What to Know Before Solving

General concepts:

* At most four intervals may be chosen
* Two intervals overlap if they share any point, including a shared boundary
* Weighted interval scheduling uses sorting by end time and a predecessor pointer
* When scores tie, the answer must be the lexicographically smallest index array
* With only four picks allowed, DP can be run separately for `k = 1..4`

Go concepts:

* Encoding interval fields into `uint64` rows for sorting and DP
* Binary search to find the latest non-overlapping predecessor
* Packing up to four 1-based indices into one `uint64` path
* Rolling two DP layers with shared backing arrays
* Using `slices.Sort` on encoded order keys
* Comparing `[]int` answers with `reflect.DeepEqual`

## Problem Description

You are given a 2D integer array `intervals`, where `intervals[i] = [li, ri, weighti]`.

Choose up to 4 non-overlapping intervals to maximize the total weight. Return the lexicographically smallest array of indices that achieves the maximum score.

## Function Signature

Expected LeetCode function signature:

```go
func maximumWeight(intervals [][]int) []int {

}
```

## Examples

### Example 1

Input:

```text
intervals = [[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]
```

Output:

```text
[2,3]
```

Explanation:

```text
Choose indices 2 and 3 with total weight 5 + 3 = 8.
```

### Example 2

Input:

```text
intervals = [[5,8,1],[6,7,7],[4,7,3],[9,10,6],[7,8,2],[11,14,3],[3,5,5]]
```

Output:

```text
[1,3,5,6]
```

Explanation:

```text
Choose indices 1, 3, 5, and 6 with total weight 7 + 6 + 3 + 5 = 21.
```

## Constraints

```text
1 <= intervals.length <= 5 * 10^4
intervals[i].length == 3
intervals[i] = [li, ri, weighti]
1 <= li <= ri <= 10^9
1 <= weighti <= 10^9
```

## Approach

This is weighted interval scheduling with an extra cap of four chosen intervals.

Steps:

1. Sort intervals by end time.
2. For each interval, binary search the latest previous interval that ends before its start.
3. Run DP for `k = 1, 2, 3, 4`:
   * `dp[i]` = best score using the first `i` sorted intervals with at most `k` picks
   * Either skip interval `i`, or take it plus the best `(k-1)`-pick solution ending at its predecessor
4. Track the best score across all `k <= 4`.
5. When scores tie, keep the lexicographically smaller index path.

To avoid storing full index arrays in DP, pack the chosen 1-based indices into a `uint64` and update it with a sorted insert helper.

## Algorithm

1. Encode each interval as:
   * `data[i] = (left << 32) | weight`
   * `order[i] = (right << 32) | index`
2. Sort `order` by end time.
3. For each sorted position, binary search predecessor `p`.
4. Rewrite each row as `(weight << 32) | predecessor | index`.
5. For each `k` from `1` to `4`:
   * Fill rolling score/path arrays over sorted intervals.
   * Transition: skip, or take with predecessor from layer `k-1`.
   * Compare by score first, then packed path for lex order.
6. Decode the best packed path back into 0-based indices.

## Why This Works

### Why sorting by end time helps

In weighted interval scheduling, once intervals are sorted by finish time, each interval has a well-defined latest compatible predecessor among earlier intervals. That makes the DP transition local.

### Why four separate layers are enough

The problem limits the answer to at most four intervals. Running DP for each exact pick count `k` and taking the best overall score is equivalent to allowing up to four picks.

### Why packed paths preserve lex order

Indices are inserted in sorted order inside the `uint64`. Comparing packed paths as integers matches lexicographic comparison of the chosen index arrays.

## Complexity Analysis

Let:

```text
n = len(intervals)
```

### Time Complexity

```text
O(n log n + n * 4)
```

Sorting costs `O(n log n)`, and each of the four DP layers scans `n` intervals with `O(1)` transitions after predecessor preprocessing.

### Space Complexity

```text
O(n)
```

Arrays for encoded intervals, order rows, and rolling DP layers all use linear space.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Treat touching intervals as overlapping because the problem counts shared boundaries as overlap
* Store predecessor index in the middle 16 bits of each encoded row
* Use `insert3414` to keep packed paths sorted when adding a new 1-based index
* Compare `q < bestP` when scores tie to enforce lexicographic order

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Single interval input
* Overlapping intervals where only one can be chosen
* Touching boundaries counted as overlap
* Three disjoint intervals all selected

## Edge Cases

Important cases to consider:

* Only one interval in the input
* Many overlapping intervals with one heavy winner
* Shared left or right boundaries
* Best answer uses fewer than four intervals
* Tie-breaking on lexicographically smaller indices

## Notes

* A standard weighted interval scheduling solution finds one best set; this problem adds both a four-interval cap and lexicographic tie-breaking.
* Packing indices into `uint64` is safe here because at most four indices are stored and each index fits in 16 bits under the given constraints.
* The predecessor binary search must use strict `<` on start/end to respect the non-overlap rule.
