---
id: 1563
title: "Stone Game V"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/stone-game-v/"
contest: "Weekly Contest 203"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Math"
  - "Dynamic Programming"
  - "Game Theory"
go_concepts:
  - "Prefix sums"
  - "Interval dynamic programming"
  - "Monotonic crossing pointer"
  - "Prefix and suffix maxima"
  - "Flattened two-dimensional tables"
  - "Reusable one-dimensional buffers"
  - "Table-driven tests"
  - "Differential tests against a cubic reference"
tags:
  - leetcode
  - go
  - array
  - math
  - dynamic-programming
  - game-theory
  - prefix-sum
  - weekly-contest-203
---

# 1563. Stone Game V

## Problem Link

LeetCode: `https://leetcode.com/problems/stone-game-v/`

## Difficulty

Hard

## Problem Topics

* Array
* Math
* Dynamic Programming
* Game Theory

## What to Know Before Solving

General concepts:

* Interval DP on a contiguous row of stones
* Prefix sums for constant-time range sums
* Bob always discards the strictly larger half
* Positive stone values make the left/right comparison monotonic
* Cached prefix and suffix maxima avoid enumerating every split

Go concepts:

* Prefix arrays of length `n + 1`
* A flattened `n * n` table instead of `[][]int`
* One reusable `[]int` of length `n` for the current left boundary
* A monotonic pointer that only moves right
* Table-driven tests plus a cubic reference implementation

## Problem Description

There are several stones arranged in a row, and each stone has an associated positive integer value given by `stoneValue`.

In each round:

1. Alice divides the current row into two non-empty contiguous rows.
2. Bob calculates the sum of each row.
3. Bob discards the row with the larger sum.
4. Alice gains points equal to the sum of the remaining row.
5. If both sums are equal, Alice decides which row is discarded.
6. The game continues using the remaining row.

The game ends when only one stone remains.

Return the maximum score Alice can obtain when playing optimally.

## Function Signature

Expected LeetCode function signature:

```go
func stoneGameV(stoneValue []int) int {

}
```

## Examples

### Example 1

Input:

```text
stoneValue = [6,2,3,4,5,5]
```

Output:

```text
18
```

### Example 2

Input:

```text
stoneValue = [7,7,7,7,7,7,7]
```

Output:

```text
28
```

### Example 3

Input:

```text
stoneValue = [4]
```

Output:

```text
0
```

A single stone cannot be split, so Alice scores nothing.

## Constraints

```text
1 <= stoneValue.length <= 500
1 <= stoneValue[i] <= 10^6
```

## Approach

Let `dp[i][j]` be the maximum score Alice can obtain from `stoneValue[i...j]`.

If the interval has fewer than two stones, `dp[i][j] = 0`.

For a split after index `k`:

```text
left  = sum(i, k)
right = sum(k+1, j)
```

The transition is:

```text
left < right:
    left + dp[i][k]

left > right:
    right + dp[k+1][j]

left == right:
    left + max(dp[i][k], dp[k+1][j])
```

Checking every split `k` for every interval `[i, j]` is the standard interval DP and costs `O(n³)`. With `n <= 500` that is too slow.

The optimized solution keeps the same recurrence and evaluates it in `O(1)` amortized time per interval.

## Monotonic Crossing Point

Every stone value is positive. For a fixed left endpoint `i`, as the split `k` moves right:

* `left` strictly increases
* `right` strictly decreases

So there is a single crossing point where the comparison changes from `left < right` to `left >= right`.

For a fixed `i`, this crossing point never moves left as `j` increases. A pointer `mid` can therefore scan all crossing points for that row in total `O(n)` time.

`mid` is the first index in `[i, j]` such that:

```text
2 * sum(i, mid) >= sum(i, j)
```

That is the first split whose left sum is at least the right sum. If every split still has `left < right`, then `mid = j`.

## Cached Maxima

Define, for the current left boundary `i`:

```text
leftBest[t] =
    max(sum(i, k) + dp[i][k])
    for i <= k <= t
```

and globally:

```text
rightBest[p][j] =
    max(sum(k, j) + dp[k][j])
    for p <= k <= j
```

These tables store exactly the quantities used by the recurrence:

* if the left part survives, Alice scores `left + dp[i][k]`
* if the right part survives, Alice scores `right + dp[k+1][j]`

At the crossing point `mid` for interval `[i, j]`:

* splits with `k < mid` have `left < right` and are answered by `leftBest[mid-1]`
* splits with `k >= mid` and a surviving right part are answered by `rightBest[mid+1][j]`
* if `sum(i, mid) == sum(mid+1, j)`, both `leftBest[mid]` and `rightBest[mid+1][j]` must be considered

A separate `dp[n][n]` matrix is unnecessary. The interval score `dp[i][j]` is computed as `score`, then immediately folded into the maxima:

```text
sum(i, j) + dp[i][j]
```

That value becomes the new candidate for `leftBest[j]` and `rightBest[i][j]`.

## Algorithm

1. If `n < 2`, return `0`.
2. Build a prefix-sum array so any range sum is `O(1)`.
3. Allocate:
   * `rightBest` of length `n * n`
   * `leftBest` of length `n`, reused for each left boundary
4. Iterate `i` from `n - 1` down to `0` so every right subinterval is already solved.
5. Initialize the singleton:
   * `leftBest[i] = stoneValue[i]`
   * `rightBest[i][i] = stoneValue[i]`
6. For `j` from `i + 1` to `n - 1`:
   * advance `mid` until `2 * sum(i, mid) >= sum(i, j)`
   * take the best surviving-left score from `leftBest`
   * take the best surviving-right score from `rightBest`
   * if the two halves are equal at `mid`, consider both
   * write `sum(i, j) + score` into both maxima tables
7. When `[i, j]` is the full array, store `score` as the answer.

## Why This Works

### Correctness of the recurrence

Alice chooses the split. Bob then discards the strictly heavier side, so Alice is forced onto the lighter side and immediately scores that side's sum. The rest of her score is the optimal continuation on the surviving interval. If the sums are equal, she may choose either continuation.

### Why the crossing pointer is enough

Because `left` increases and `right` decreases, the set of splits with `left < right` is always a prefix of `[i, j)`. The maximum of a prefix of `left + dp[i][k]` values is exactly `leftBest`. The complementary suffix is exactly `rightBest`. No other split class exists.

### Why `mid` is monotonic in `j`

Increasing `j` appends a positive stone, so the total sum grows and the inequality `2 * sum(i, mid) >= total` can only become harder. Therefore `mid` never decreases.

### Why a `dp` matrix is not stored

After `score = dp[i][j]` is known, later intervals only need maxima of `sum + dp`, not the raw `dp` values. Those maxima are exactly `leftBest` and `rightBest`.

## Complexity Analysis

Let:

```text
n = len(stoneValue)
```

### Time Complexity

```text
O(n²)
```

There are `O(n²)` intervals. For each fixed `i`, the pointer `mid` moves at most `n` times, and every `j` does `O(1)` extra work after that. This is the optimized standard solution for this problem.

### Space Complexity

```text
O(n²)
```

`rightBest` uses one flattened table of size `n²`. `leftBest` and the prefix array are `O(n)`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Range sums use `prefix[j+1] - prefix[i]`
* `rightBest[p][j]` is stored at `rightBest[p*n+j]`
* `leftBest` is overwritten for each left boundary `i`
* Intervals are processed with decreasing `i` so `rightBest[i+1][j]` already exists
* The full-array answer is `score`, not `total + score`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Two-stone cases with equal, left-smaller, and right-smaller values
* A three-stone interval that needs an equal split
* Randomized differential tests against an `O(n³)` reference

## Edge Cases

Important cases to consider:

* `n = 1`, which cannot be split
* `n = 2`, where Alice scores the smaller stone
* All values equal
* A split where the two halves have equal sums
* A crossing point at the first or last legal split

## Notes

* The cubic interval DP is useful as a reference, but not as the submitted solution for `n <= 500`.
* Memoized recursion with pruning can pass, but the monotonic-maxima form is cleaner and worst-case `O(n²)`.
* Do not interpret `O(n²)` as a proven lower bound; it is the optimized standard algorithm for this problem.
