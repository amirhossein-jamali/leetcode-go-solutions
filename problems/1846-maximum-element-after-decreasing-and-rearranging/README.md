---
id: 1846
title: "Maximum Element After Decreasing and Rearranging"
difficulty: "Medium"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/"
contest: "Biweekly Contest 51"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Greedy"
  - "Sorting"
go_concepts:
  - "In-place counting sort with encoded frequencies"
  - "Capping values at array length"
  - "Greedy construction of a valid non-decreasing chain"
  - "Table-driven tests with slice copies"
tags:
  - leetcode
  - go
  - array
  - greedy
  - counting-sort
  - biweekly-contest-51
---

# 1846. Maximum Element After Decreasing and Rearranging

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/`

## Difficulty

Medium

## Problem Topics

* Array
* Greedy
* Sorting

## What to Know Before Solving

General concepts:

* You may rearrange the array in any order
* You may only decrease values, never increase them
* The final array must start at `1`
* Adjacent values can differ by at most `1`
* The goal is to maximize the largest final value

Go concepts:

* Counting frequencies without a separate hash map
* Encoding counts inside the input slice using multiples of `n`
* Capping large values before counting
* Greedy accumulation with a running maximum cap
* Copying input slices in tests because the solution mutates `arr`

## Problem Description

You are given an array of positive integers `arr`. Perform some operations (possibly none) on `arr` so that it satisfies these conditions:

* The value of the first element in `arr` must be `1`.
* The absolute difference between any two adjacent elements must be less than or equal to `1`.

There are two types of operations that you can perform any number of times:

* Decrease the value of any element of `arr` to a smaller positive integer.
* Rearrange the elements of `arr` to be in any order.

Return the maximum possible value of an element in `arr` after performing the operations to satisfy the conditions.

## Function Signature

Expected LeetCode function signature:

```go
func maximumElementAfterDecrementingAndRearranging(arr []int) int {

}
```

## Examples

### Example 1

Input:

```text
arr = [2,2,1,2,1]
```

Output:

```text
2
```

Explanation:

```text
We can rearrange arr to [1,2,2,2,1].
The largest element is 2.
```

### Example 2

Input:

```text
arr = [100,1,1000]
```

Output:

```text
3
```

Explanation:

```text
One valid sequence of operations is:

1. Rearrange to [1,100,1000]
2. Decrease 100 to 2
3. Decrease 1000 to 3

Now arr = [1,2,3], and the largest element is 3.
```

### Example 3

Input:

```text
arr = [1,2,3,4,5]
```

Output:

```text
5
```

Explanation:

```text
The array already satisfies the conditions.
The largest element is 5.
```

## Constraints

```text
1 <= arr.length <= 10^5
1 <= arr[i] <= 10^9
```

## Approach

After rearranging, the best valid array is non-decreasing and increases by at most `1` each step, starting from `1`.

So the ideal shape looks like:

```text
1, 2, 2, 3, 3, 3, ...
```

If we know how many numbers can be assigned to each value, we can greedily build the longest valid chain.

Key observations:

* Any value greater than `n` can be decreased to at most `n`, because we only have `n` positions.
* We only need frequencies of values `1..n`.
* A counting-sort style pass can store frequencies inside the same array.

Then scan values from `1` to `n`:

* Add the frequency of the current value to a running total
* The running total means how many elements can reach at least this value
* Cap the running total at the current value itself

The final capped maximum is the answer.

## Algorithm

1. Let `n = len(arr)`.
2. Cap every value above `n` down to `n`.
3. Use in-place counting sort to store frequencies of values `1..n`.
4. Initialize `maxValue = 0`.
5. For each `value` from `1` to `n`:
   * Read its frequency
   * Add it to `maxValue`
   * Cap `maxValue` at `value`
6. Return `maxValue`.

## Why This Works

Because rearranging is free, only frequencies matter.

To build a valid array starting at `1`, each next value can stay the same or increase by `1`. That means if we want the final maximum to be `M`, we need enough elements to fill the tiers:

```text
1, then 2, then 3, ..., then M
```

Scanning from small to large and accumulating frequencies tells us how many elements can support each level.

If the accumulated count exceeds the current value, we cannot assign more elements to higher tiers without breaking the `+1` rule, so we cap at the current value.

Therefore, the last capped value is the maximum achievable top element.

## Complexity Analysis

Let `n` be the length of `arr`.

### Time Complexity

```text
O(n)
```

Capping, counting, and scanning all take linear time.

### Space Complexity

```text
O(1)
```

The counting sort is done in place on the input slice.

Note: the solution mutates `arr`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Values above `n` are capped before counting
* Frequencies are encoded by adding multiples of `n`
* The greedy cap `maxValue = min(maxValue, value)` is the core logic

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Single-element input
* Repeated large values
* Two-element input

## Edge Cases

Important cases to consider:

* `arr` already sorted and valid
* all elements equal and large
* one element only
* many duplicates of the same small value
* large values that must be capped to `n`

## Notes

* This problem is greedy plus counting sort, not full sorting of all values.
* Rearranging makes the original order irrelevant; only frequencies matter.
* Copy `arr` in local tests because the function modifies the input slice.
