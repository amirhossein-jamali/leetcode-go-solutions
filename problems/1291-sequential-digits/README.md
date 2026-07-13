---
id: 1291
title: "Sequential Digits"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/sequential-digits/"
contest: "Weekly Contest 167"
status: "Solved"
language: "Go"
topics:
  - "Enumeration"
  - "Math"
go_concepts:
  - "Preallocated slices with make"
  - "Nested loops over fixed small search space"
  - "Arithmetic progression generation"
  - "Early return when values exceed high"
  - "Table-driven tests with reflect.DeepEqual"
tags:
  - leetcode
  - go
  - enumeration
  - math
  - weekly-contest-167
---

# 1291. Sequential Digits

## Problem Link

LeetCode: `https://leetcode.com/problems/sequential-digits/`

## Difficulty

Medium

## Problem Topics

* Enumeration
* Math

## What to Know Before Solving

General concepts:

* Sequential digits mean each digit is exactly one more than the previous digit
* Valid numbers always start with digits `1` through `8`
* The maximum length is `9` because the sequence cannot continue past digit `9`
* There are only `36` sequential-digit integers in total, so brute enumeration is efficient
* For a fixed length, all valid numbers form an arithmetic progression

Go concepts:

* Preallocating a result slice with `make([]int, 0, 36)`
* Nested loops over a tiny fixed search space
* Updating `first` and `step` when moving to the next digit length
* Early return once generated values exceed `high`
* Comparing `[]int` results in tests with `reflect.DeepEqual`

## Problem Description

An integer has sequential digits if and only if each digit in the number is one more than the previous digit.

Return a sorted list of all integers in the range `[low, high]` inclusive that have sequential digits.

## Function Signature

Expected LeetCode function signature:

```go
func sequentialDigits(low int, high int) []int {

}
```

## Examples

### Example 1

Input:

```text
low = 100, high = 300
```

Output:

```text
[123,234]
```

### Example 2

Input:

```text
low = 1000, high = 13000
```

Output:

```text
[1234,2345,3456,4567,5678,6789,12345]
```

## Constraints

```text
10 <= low <= high <= 10^9
```

## Approach

Enumerate all sequential-digit numbers directly instead of checking every integer in `[low, high]`.

For length `2`, the numbers are `12, 23, 34, ..., 89` with step `11`.

For length `3`, they are `123, 234, ..., 789` with step `111`.

In general, for length `L`:

* the first number is `12...L` in sequential form
* the step is `11...1` with `L - 1` ones
* there are `10 - L` valid numbers

Iterate lengths from `2` to `9`, generate each arithmetic progression, keep values inside `[low, high]`, and stop early when a value exceeds `high`.

## Algorithm

1. Initialize `result`, `first = 12`, and `step = 11`.
2. For each digit length `digits` from `2` to `9`:
   * Generate `10 - digits` numbers starting at `first` with step `step`.
   * If `current > high`, return `result` immediately.
   * If `current >= low`, append `current` to `result`.
   * Update `first = first*10 + digits + 1` and `step = step*10 + 1`.
3. Return `result`.

## Why This Works

### Why enumeration is enough

The search space is tiny. Even over the full constraint range, there are only `36` sequential-digit numbers:

```text
8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 = 36
```

So checking every integer up to `10^9` is unnecessary.

### Why the arithmetic progression is correct

For a fixed length `L`, the first digit can only be `1` through `9 - L + 1`.

Once the first digit is chosen, every later digit is forced. That gives exactly one number per starting digit, and adjacent numbers differ by `111...1`, so they form an arithmetic progression.

### Why the output is sorted

We iterate lengths in increasing order, and within each length we generate values in increasing order. Therefore the appended results are already sorted.

## Complexity Analysis

There are at most `36` sequential-digit numbers total.

### Time Complexity

```text
O(1)
```

The number of generated values is bounded by a constant.

### Space Complexity

```text
O(1)
```

aside from the output slice, which contains at most `36` integers.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Preallocate `result` with capacity `36` because that is the maximum possible answer size.
* Use early return when `current > high` because later values in the same length and all longer lengths are larger.
* Update `first` and `step` with `first*10 + digits + 1` and `step*10 + 1` to move to the next digit length.

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Full two-digit sequential range
* Single-value range
* Range with no valid answer
* Range ending before the first match
* The longest sequential number `123456789`

## Edge Cases

Important cases to consider:

* `low` and `high` with no sequential-digit number between them
* Range containing exactly one valid number
* Two-digit answers such as `12`
* Nine-digit answer `123456789`
* `high` smaller than the first candidate for a given length

## Notes

* BFS or DFS that appends the next digit also works, but the arithmetic-progression form is shorter.
* Because the valid set is so small, this problem is essentially fixed enumeration rather than range scanning.
* The result is naturally sorted, so no extra sorting step is needed.
