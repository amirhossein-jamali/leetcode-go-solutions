---
id: 115
title: "Distinct Subsequences"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/distinct-subsequences/"
status: "Solved"
language: "Go"
topics:
  - "String"
  - "Dynamic Programming"
go_concepts:
  - "1D DP with space optimization"
  - "Backward inner loop to avoid overwriting"
  - "String indexing by byte"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - string
  - dynamic-programming
---

# 115. Distinct Subsequences

## Problem Link

LeetCode: `https://leetcode.com/problems/distinct-subsequences/`

## Difficulty

Hard

## Problem Topics

* String
* Dynamic Programming

## What to Know Before Solving

General concepts:

* A subsequence keeps relative order but may skip characters
* Two strings can match in many different ways when characters repeat
* `dp[j]` can represent how many ways the first `j` characters of `t` can be formed
* Updating `dp` from right to left prevents using the same character twice in one step

Go concepts:

* Allocating a DP slice with `make([]int, m+1)`
* Iterating `j` from `m-1` down to `0`
* Comparing characters with `s[i] == t[j]`
* Accumulating counts with `dp[j+1] += dp[j]`
* Writing table-driven tests with the `testing` package

## Problem Description

Given two strings `s` and `t`, return the number of distinct subsequences of `s` which equal `t`.

Since the answer may be large, the result is guaranteed to fit in a 32-bit signed integer.

## Function Signature

Expected LeetCode function signature:

```go
func numDistinct(s string, t string) int {

}
```

## Examples

### Example 1

Input:

```text
s = "rabbbit", t = "rabbit"
```

Output:

```text
3
```

Explanation:

```text
There are 3 ways to form "rabbit" from "rabbbit" by choosing different positions for the repeated 'b'.
```

### Example 2

Input:

```text
s = "babgbag", t = "bag"
```

Output:

```text
5
```

Explanation:

```text
There are 5 distinct ways to pick characters from "babgbag" that form "bag".
```

## Constraints

```text
1 <= s.length, t.length <= 1000
s and t consist of English letters
```

## Approach

Use dynamic programming where `dp[j]` stores the number of ways to form the first `j` characters of `t`.

Base case:

* `dp[0] = 1` because the empty prefix of `t` can always be formed in exactly one way

Transition:

* Scan `s` from left to right
* For each character `s[i]`, update `dp` from right to left
* If `s[i] == t[j]`, then every way to form prefix `t[:j]` can also extend to prefix `t[:j+1]`
* So add `dp[j]` into `dp[j+1]`

The final answer is `dp[m]`, where `m = len(t)`.

## Algorithm

1. Let `m = len(t)` and create `dp` of size `m+1`.
2. Set `dp[0] = 1`.
3. For each character `s[i]`:
   * Iterate `j` from `m-1` down to `0`.
   * If `s[i] == t[j]`, do `dp[j+1] += dp[j]`.
4. Return `dp[m]`.

## Why This Works

### Why `dp[j]` is correct

After processing some prefix of `s`, `dp[j]` counts exactly how many distinct subsequences of that processed prefix match `t[:j]`.

When a new character `s[i]` arrives:

* If it does not equal `t[j]`, nothing changes for position `j+1`
* If it equals `t[j]`, every existing way to match `t[:j]` can use this character to create a new way to match `t[:j+1]`

That is exactly `dp[j+1] += dp[j]`.

### Why the inner loop goes backward

If we updated from left to right, the same `s[i]` could contribute multiple times within one iteration.

Example:

```text
s = "a", t = "a"
```

Forward update would incorrectly turn `dp[1]` into `2`.

Backward update ensures each character of `s` is used at most once per DP layer.

## Complexity Analysis

Let:

```text
n = len(s)
m = len(t)
```

### Time Complexity

```text
O(n * m)
```

We scan all `n` characters of `s`, and for each one we update up to `m` DP states.

### Space Complexity

```text
O(m)
```

Only one DP array of length `m+1` is needed because each row depends only on the previous row.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Initialize `dp[0] = 1` for the empty target prefix
* Iterate `j` from `m-1` down to `0` to preserve 1D DP correctness
* Return `dp[m]`, which counts full matches of `t`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Exact match with no extra characters
* Target longer than source
* No valid subsequence
* Repeated characters in source
* Multiple ways to skip characters
* Minimum single-character input

## Edge Cases

Important cases to consider:

* `s == t` with exactly one match
* `len(t) > len(s)` so answer is `0`
* Repeated letters in `s` creating many distinct match paths
* Source much longer than target
* Single-character strings

## Notes

* A full 2D DP table is easier to reason about, but the 1D optimized version is standard for interviews.
* This problem is closely related to edit distance and subsequence DP templates.
* The backward update trick appears often in knapsack-style and counting DP problems.
