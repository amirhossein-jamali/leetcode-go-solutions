---
id: 3699
title: "Number of ZigZag Arrays I"
difficulty: "Hard"
level: "Senior Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/number-of-zigzag-arrays-i/"
contest: "Weekly Contest 469"
status: "Solved"
language: "Go"
topics:
  - "Dynamic Programming"
  - "Prefix Sum"
go_concepts:
  - "uint32 for modulo arithmetic"
  - "1D DP slice over value range"
  - "Prefix and suffix rolling sums"
  - "Modulo subtraction instead of overflow-prone addition"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - dynamic-programming
  - prefix-sum
  - weekly-contest-469
---

# 3699. Number of ZigZag Arrays I

## Problem Link

LeetCode: `https://leetcode.com/problems/number-of-zigzag-arrays-i/`

## Difficulty

Hard

## Problem Topics

* Dynamic Programming
* Prefix Sum

## What to Know Before Solving

General concepts:

* ZigZag arrays must alternate direction and avoid equal adjacent values
* No three consecutive elements may be strictly increasing or strictly decreasing
* Counting valid arrays over a numeric range is a DP problem
* Prefix sums can optimize transitions over all smaller or larger values
* Final answer must be taken modulo `10^9 + 7`

Go concepts:

* Using `uint32` for counts and modulo operations
* A 1D DP array indexed by offset value in `[l, r]`
* Updating DP with rolling prefix and suffix sums
* Using `if sum >= mod { sum -= mod }` to stay in range
* Multiplying the final count by 2 for both starting directions
* Writing table-driven tests with the `testing` package

## Problem Description

You are given three integers `n`, `l`, and `r`.

A ZigZag array of length `n` is defined as follows:

* Each element lies in the range `[l, r]`.
* No two adjacent elements are equal.
* No three consecutive elements form a strictly increasing or strictly decreasing sequence.

Return the total number of valid ZigZag arrays.

Since the answer may be large, return it modulo `10^9 + 7`.

## Function Signature

Expected LeetCode function signature:

```go
func zigZagArrays(n int, l int, r int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 3, l = 4, r = 5
```

Output:

```text
2
```

Explanation:

```text
There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:

[4, 5, 4]
[5, 4, 5]
```

### Example 2

Input:

```text
n = 3, l = 1, r = 3
```

Output:

```text
10
```

Explanation:

```text
There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

[1, 2, 1], [1, 3, 1], [1, 3, 2]
[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
[3, 1, 2], [3, 1, 3], [3, 2, 3]
```

## Constraints

```text
3 <= n <= 2000
1 <= l < r <= 2000
```

## Approach

Map values from `[l, r]` to indices `0..m-1` where `m = r - l + 1`.

Use DP where `dp[v]` counts valid partial arrays ending at value `v`.

The solution builds arrays in two-step chunks:

1. A prefix-sum pass models adding the next element that must be smaller than the previous one.
2. A suffix-sum pass models adding the next element that must be larger than the previous one.

Start with length 1 arrays: every value has exactly one array.

Repeat until length reaches `n`. The loop advances by 2 because each full iteration adds two positions to the pattern.

Finally, sum all ending states and multiply by 2 to account for both possible starting directions.

## Algorithm

1. Set `m = r - l + 1`.
2. Initialize `dp[v] = 1` for all values.
3. For `length = 2, 4, 6, ...` up to `n`:
   * Prefix pass left to right:
     * Replace `dp[v]` with the sum of previous states with smaller values.
   * If `length == n`, stop.
   * Suffix pass right to left:
     * Replace `dp[v]` with the sum of previous states with larger values.
4. Sum all values in `dp`.
5. Return `(sum * 2) % mod`.

## Why This Works

At any step, the next value must differ from the previous one and must not create a 3-length monotone segment.

The prefix pass counts ways to extend with a smaller next value. Because all smaller values contribute, a rolling prefix sum gives the transition in one scan.

The suffix pass does the symmetric work for larger next values.

Processing two positions per outer loop matches the alternating up/down structure of a ZigZag array. Initializing every value with one array of length 1 covers all possible starting points. Multiplying by 2 at the end accounts for both upward-first and downward-first patterns.

## Complexity Analysis

Let `m = r - l + 1`.

### Time Complexity

```text
O(n * m)
```

Each of the `O(n)` DP layers performs `O(m)` prefix or suffix work.

### Space Complexity

```text
O(m)
```

Only one DP array of size `m` is stored.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Values are shifted to zero-based indices with `m = r - l + 1`.
* Modulo is kept inline with subtraction to avoid overflow.
* The outer loop stops early when `length == n` after the prefix pass if `n` is odd in the two-step cycle.

```go
package numberofzigzagarraysi

const mod uint32 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1

	dp := make([]uint32, m)

	for i := range dp {
		dp[i] = 1
	}

	for length := 2; length <= n; length += 2 {
		var prefix uint32

		for value := 0; value < m; value++ {
			old := dp[value]
			dp[value] = prefix

			prefix += old
			if prefix >= mod {
				prefix -= mod
			}
		}

		if length == n {
			break
		}

		var suffix uint32

		for value := m - 1; value >= 0; value-- {
			old := dp[value]
			dp[value] = suffix

			suffix += old
			if suffix >= mod {
				suffix -= mod
			}
		}
	}

	var onePattern uint32

	for _, count := range dp {
		onePattern += count

		if onePattern >= mod {
			onePattern -= mod
		}
	}

	return int(uint64(onePattern) * 2 % uint64(mod))
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Minimum two-value range
* A longer array length (`n = 5`)

## Edge Cases

Important cases to consider:

* Smallest valid `n = 3`
* Range with only two distinct values
* Larger `n` with small value range
* Large counts requiring modulo arithmetic
* Even and odd handling in the two-step DP loop

## Notes

* This is a counting DP problem, not a construction problem.
* Prefix and suffix rolling sums are the key optimization over naive transitions.
* The final multiplication by 2 is easy to forget, but it is required for both ZigZag directions.
