---
id: 3739
title: "Count Subarrays With Majority Element II"
difficulty: "Hard"
level: "Senior Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/count-subarrays-with-majority-element-ii/"
contest: "Biweekly Contest 169"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
  - "Prefix Sum"
go_concepts:
  - "int64 for large subarray counts"
  - "In-place array transformation"
  - "Prefix balance tracking"
  - "Frequency counting with encoded array slots"
  - "Table-driven tests with slice copies"
tags:
  - leetcode
  - go
  - array
  - hash-table
  - prefix-sum
  - biweekly-contest-169
---

# 3739. Count Subarrays With Majority Element II

## Problem Link

LeetCode: `https://leetcode.com/problems/count-subarrays-with-majority-element-ii/`

## Difficulty

Hard

## Problem Topics

* Array
* Hash Table
* Prefix Sum

LeetCode also tags alternative approaches such as divide and conquer, segment tree, and merge sort. This solution uses the same linear in-place prefix-counting method as part I.

## What to Know Before Solving

General concepts:

* Same rules as [3737. Count Subarrays With Majority Element I](./../3737-count-subarrays-with-majority-element-i/)
* The answer can be very large, so the return type is `int64`
* `nums.length` can be up to `10^5`, so an `O(n^2)` brute force is too slow
* Transforming `target` to `1` and other values to `0` reduces the problem to prefix-balance counting
* Prefix frequencies can be stored inside the transformed array using base encoding

Go concepts:

* Returning `int64` instead of `int`
* Accumulating the answer with `answer += int64(smaller)`
* Rewriting slice values in place
* Using `base = n + 1` to pack counts into array slots
* Copying input slices in tests because the solution mutates `nums`

## Problem Description

You are given an integer array `nums` and an integer `target`.

Return the number of subarrays of `nums` in which `target` is the majority element.

The majority element of a subarray is the element that appears strictly more than half of the times in that subarray.

This is the large-`n` version of part I.

## Function Signature

Expected LeetCode function signature:

```go
func countMajoritySubarrays(nums []int, target int) int64 {

}
```

## Examples

### Example 1

Input:

```text
nums = [1,2,2,3], target = 2
```

Output:

```text
5
```

Explanation:

```text
Valid subarrays with target = 2 as the majority element:

nums[1..1] = [2]
nums[2..2] = [2]
nums[1..2] = [2,2]
nums[0..2] = [1,2,2]
nums[1..3] = [2,2,3]
```

### Example 2

Input:

```text
nums = [1,1,1,1], target = 1
```

Output:

```text
10
```

Explanation:

```text
All 10 subarrays have 1 as the majority element.
```

### Example 3

Input:

```text
nums = [1,2,3], target = 4
```

Output:

```text
0
```

Explanation:

```text
target = 4 does not appear in nums at all.
Therefore, there cannot be any subarray where 4 is the majority element.
```

## Constraints

```text
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= target <= 10^9
```

## Approach

The logic is the same as part I:

1. Rewrite `nums` so `target` becomes `1` and every other value becomes `0`.
2. Track a prefix balance where each `1` increases the balance and each `0` decreases it.
3. Count how many previous prefix states make the current subarray a majority subarray.
4. Store prefix frequencies in the same array using:

```text
stored value = actual value + count * base
```

where `base = n + 1`.

The only practical differences for part II are:

* `n` can be much larger
* the answer must be accumulated in `int64`

## Algorithm

1. Transform `nums` into `1`/`0` and count non-target elements as `otherCount`.
2. Initialize encoded frequency storage with `base = n + 1`.
3. Scan once from left to right:
   * Update `prefix` and `smaller` based on whether the current value is target.
   * Add `smaller` to `answer`.
   * Update the encoded frequency table for the current prefix index.
4. Return `answer`.

## Why This Works

After transformation, a subarray has `target` as majority exactly when it contains more `1`s than `0`s.

While scanning, each index contributes the number of valid starting positions for a majority subarray ending at that index. The encoded frequency table lets us query and update prefix counts in constant time per step.

Because each index is processed once, the algorithm remains linear even for `n = 10^5`.

## Complexity Analysis

Let `n` be the length of `nums`.

### Time Complexity

```text
O(n)
```

### Space Complexity

```text
O(1)
```

Aside from the input slice, only a few integer variables are used.

Note: the solution mutates `nums` in place.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* The return type is `int64` for large answers
* `answer += int64(smaller)` avoids overflow during accumulation
* The core logic matches part I exactly
* `base = n + 1` keeps encoded counts separate from transformed values

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Single-element arrays
* Arrays with no target occurrences
* A case where only length-one subarrays can have majority

## Edge Cases

Important cases to consider:

* `target` never appears in `nums`
* all elements equal `target`
* only one element
* very large `n` with sparse target occurrences
* answer larger than 32-bit integer range

## Notes

* Part II is mainly a scalability version of part I.
* LeetCode tags heavier tools like segment trees, but the same linear encoded-prefix approach works here too.
* Always copy `nums` in local tests because the function modifies the input slice.

## Related Problem

* [3737. Count Subarrays With Majority Element I](./../3737-count-subarrays-with-majority-element-i/) — same rules, but `nums.length <= 1000` and return type `int`.
