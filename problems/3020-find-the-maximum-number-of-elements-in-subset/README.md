---
id: 3020
title: "Find the Maximum Number of Elements in Subset"
difficulty: "Medium"
level: "Senior"
platform: "LeetCode"
link: "https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/"
contest: "Weekly Contest 382"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
  - "Enumeration"
go_concepts:
  - "map[int64]int for frequency counting"
  - "int64 to avoid overflow when squaring"
  - "Enumerating chain starts from frequency map"
  - "Special handling for value 1"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - hash-table
  - enumeration
  - weekly-contest-382
---

# 3020. Find the Maximum Number of Elements in Subset

## Problem Link

LeetCode: `https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/`

## Difficulty

Medium

## Problem Topics

* Array
* Hash Table
* Enumeration

## What to Know Before Solving

General concepts:

* The required pattern is a palindrome of powers: `[x, x^2, x^4, ..., x^{k/2}, x^k, x^{k/2}, ..., x^4, x^2, x]`
* Each power appears twice except the peak `x^k`, which appears once
* Extending a chain adds two elements at a time
* Value `1` is special because every power of `1` is still `1`
* Frequency counts determine how far a chain can grow

Go concepts:

* Building a frequency map with `map[int64]int`
* Casting `int` to `int64` before squaring to reduce overflow risk
* Iterating over map entries with `for start, count := range freq`
* Handling the special case for `1` separately
* Writing table-driven tests with the `testing` package

## Problem Description

You are given an array of positive integers `nums`.

You need to select a subset of `nums` which satisfies the following condition:

You can place the selected elements in a `0`-indexed array such that it follows the pattern:

```text
[x, x^2, x^4, ..., x^{k/2}, x^k, x^{k/2}, ..., x^4, x^2, x]
```

Note that `k` can be any non-negative power of `2`.

For example:

* `[2, 4, 16, 4, 2]` follows the pattern
* `[3, 9, 3]` follows the pattern
* `[2, 4, 8, 4, 2]` does not

Return the maximum number of elements in a subset that satisfies these conditions.

## Function Signature

Expected LeetCode function signature:

```go
func maximumLength(nums []int) int {

}
```

## Examples

### Example 1

Input:

```text
nums = [5,4,1,2,2]
```

Output:

```text
3
```

Explanation:

```text
We can select the subset {4,2,2}, which can be placed as [2,4,2].
Since 2^2 == 4, the answer is 3.
```

### Example 2

Input:

```text
nums = [1,3,2,4]
```

Output:

```text
1
```

Explanation:

```text
We can select the subset {1}, which can be placed as [1].
Other singleton subsets such as {2}, {3}, or {4} also work.
```

## Constraints

```text
2 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
```

## Approach

Build a frequency map of all values in `nums`.

Every valid subset corresponds to a chain:

```text
x -> x^2 -> x^4 -> ... -> x^k
```

In the palindrome pattern, each intermediate power is used twice, except the peak power which is used once.

So if we can extend a chain while each required power appears at least twice, the subset length grows by `2` each time.

### Special case for `1`

Every power of `1` is still `1`. If there are `ones` copies of `1`, the best odd-length subset of all ones has size:

```text
ones - 1 + ones % 2
```

### General chains

For every starting value `x != 1` with frequency at least `2`:

1. Start with length `1` using the peak conceptually at `x`.
2. Try to move to `x^2`, then `(x^2)^2`, and so on.
3. Each successful step adds `2` elements to the subset length.
4. Stop when the next square is missing from the frequency map or its count drops below `2`.

Track the maximum length across all starts.

## Algorithm

1. Count frequencies in `freq`.
2. Initialize `answer = 1`.
3. If `1` exists, update `answer` using the special formula for ones.
4. For each `(start, count)` in `freq`:
   * Skip if `start == 1` or `count < 2`.
   * Set `length = 1` and `current = start`.
   * While `count >= 2`:
     * Compute `next = current * current`.
     * If `freq[next] == 0`, stop.
     * Add `2` to `length`.
     * Move `current = next` and `count = freq[next]`.
   * Update `answer`.
5. Return `answer`.

## Why This Works

The pattern always uses powers of a single base `x` in symmetric order.

Except for the peak, every power appears twice. That means whenever we extend from `x^a` to `x^{2a}`, we need at least two copies of `x^{2a}` in the array.

Starting from each possible base and greedily extending while frequencies allow it explores every valid chain shape.

The special formula for `1` handles the fact that squaring does not change the value, so the usual chain logic does not apply.

Therefore, the maximum over all starts is the answer.

## Complexity Analysis

Let `n` be the length of `nums` and let `U` be the number of distinct values.

### Time Complexity

```text
O(n + U log max(nums))
```

Building the frequency map takes `O(n)`.

Each chain squares its current value until it stops. The number of steps per start is bounded by how many times you can square before exceeding `10^9`.

### Space Complexity

```text
O(U)
```

The frequency map stores distinct values.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Frequencies use `int64` keys because squaring can exceed 32-bit range
* Chains only start from values with frequency at least `2`
* Value `1` is handled separately before the general loop

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Three copies of `1`
* A valid `2, 2, 4` chain
* A longer chain with repeated powers

## Edge Cases

Important cases to consider:

* Only singleton values available
* Multiple copies of `1`
* Chains that stop because the next square is missing
* Chains that stop because the next power appears only once
* Starting values greater than `1` with exactly two copies

## Notes

* This is an enumeration plus frequency-check problem, not a full subset search.
* Squaring repeatedly quickly exceeds `10^9`, so chains are short in practice.
* Always handle `1` separately; treating it like other bases is incorrect.
