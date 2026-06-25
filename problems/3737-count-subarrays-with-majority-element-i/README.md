---
id: 3737
title: "Count Subarrays With Majority Element I"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/count-subarrays-with-majority-element-i/"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Prefix Sum"
  - "Hash Table"
go_concepts:
  - "In-place array transformation"
  - "Prefix balance tracking"
  - "Frequency counting with encoded array slots"
  - "Integer division and modulo for packed counts"
  - "Table-driven tests with slice copies"
tags:
  - leetcode
  - go
  - array
  - prefix-sum
  - hash-table
---

# 3737. Count Subarrays With Majority Element I

## Problem Link

LeetCode: `https://leetcode.com/problems/count-subarrays-with-majority-element-i/`

## Difficulty

Medium

## Problem Topics

* Array
* Prefix Sum
* Hash Table

## What to Know Before Solving

General concepts:

* A majority element appears strictly more than half the times in a subarray
* Comparing counts of `target` vs non-target values is enough
* Prefix sums over a transformed array can track balance between target and other elements
* Subarray counting often becomes counting valid prefix differences

Go concepts:

* Rewriting slice values in place (`1` for target, `0` otherwise)
* Maintaining a running prefix balance while scanning
* Using an encoded frequency table inside the same slice
* Separating actual values from counts with `base = n + 1`
* Copying input slices in tests because the solution mutates `nums`
* Writing table-driven tests with the `testing` package

## Problem Description

You are given an integer array `nums` and an integer `target`.

Return the number of subarrays of `nums` in which `target` is the majority element.

The majority element of a subarray is the element that appears strictly more than half of the times in that subarray.

## Function Signature

Expected LeetCode function signature:

```go
func countMajoritySubarrays(nums []int, target int) int {

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

So there are 5 such subarrays.
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
1 <= nums.length <= 1000
1 <= nums[i] <= 10^9
1 <= target <= 10^9
```

## Approach

Transform the array so that `target` becomes `1` and every other value becomes `0`.

For a subarray, `target` is the majority element if and only if:

```text
count(target) > count(other)
```

After transformation, this is equivalent to tracking a prefix balance where `1` increases the balance and `0` decreases it.

The solution scans from left to right and maintains:

* `prefix`: current balance of target vs other elements
* `smaller`: how many previous prefix states are small enough to form valid majority subarrays ending at the current index
* a frequency table of prefix balances

Instead of a separate hash map, prefix frequencies are stored inside the transformed `nums` slice using:

```text
stored value = actual value + count * base
```

where `base = n + 1`.

## Algorithm

1. Rewrite `nums`:
   * `target` becomes `1`
   * other values become `0`
   * count how many non-target values exist as `otherCount`
2. Initialize frequency storage with `base = n + 1`.
3. Scan the array once:
   * If current value is target (`nums[i] % base == 1`), increase `prefix`.
   * Otherwise, decrease `prefix`.
   * Update `smaller` using the frequency of the relevant prefix index.
   * Add `smaller` to `answer`.
   * Update the frequency table for the current prefix index.
4. Return `answer`.

## Why This Works

After transforming values, a subarray has `target` as majority exactly when the transformed subarray contains more `1`s than `0`s.

While scanning right to left conceptually, each new position can extend previous prefix states. The `prefix` variable tracks the balance, and counting how many earlier balances are valid gives the number of majority subarrays ending at the current index.

Encoding frequencies inside the same array avoids an extra hash map and keeps the solution linear in time.

## Complexity Analysis

Let `n` be the length of `nums`.

### Time Complexity

```text
O(n)
```

The array is scanned a constant number of times.

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

* `nums` is rewritten to `1`/`0`, so the input slice is modified
* `base = n + 1` separates stored values from frequency counts
* `otherCount` shifts prefix indices for the frequency table
* `zeroFrequency` handles the special prefix index `0`

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
* subarrays of length `1` containing `target`
* subarrays of length `2` where majority is impossible unless both are `target`

## Notes

* This is a prefix-balance counting problem disguised as a majority problem.
* The in-place frequency encoding is memory-efficient but modifies the input array.
* For local tests, copy the slice before calling the function if you need to preserve the original input.
