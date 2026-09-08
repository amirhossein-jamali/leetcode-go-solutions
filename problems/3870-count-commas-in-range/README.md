---
id: 3870
title: "Count Commas in Range"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/count-commas-in-range/"
contest: "Weekly Contest 493"
status: "Solved"
language: "Go"
topics:
  - "Math"
go_concepts:
  - "Threshold checks"
  - "Direct formula instead of simulation"
  - "Inclusive counting with n - 999"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - math
  - weekly-contest-493
---

# 3870. Count Commas in Range

## Problem Link

LeetCode: `https://leetcode.com/problems/count-commas-in-range/`

## Difficulty

Easy

## Problem Topics

* Math

## What to Know Before Solving

General concepts:

* Standard number formatting inserts a comma after every three digits from the right
* Numbers with fewer than four digits contain no commas
* For `1 <= n <= 10^5`, every integer from `1000` to `n` uses exactly one comma
* Counting commas is equivalent to counting how many integers in `[1000, n]` exist

Go concepts:

* Early return for small inputs
* Using a direct arithmetic formula instead of looping through every integer
* Writing table-driven tests with the `testing` package

## Problem Description

You are given an integer `n`.

Return the total number of commas used when writing all integers from `1` to `n` inclusive in standard number formatting.

## Function Signature

Expected LeetCode function signature:

```go
func countCommas(n int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 1002
```

Output:

```text
3
```

Explanation:

```text
The numbers "1,000", "1,001", and "1,002" each contain one comma, giving a total of 3.
```

### Example 2

Input:

```text
n = 998
```

Output:

```text
0
```

Explanation:

```text
All numbers from 1 to 998 have fewer than four digits. Therefore, no commas are used.
```

## Constraints

```text
1 <= n <= 10^5
```

## Approach

Simulating every integer from `1` to `n` is unnecessary.

Under the given constraints:

* If `n < 1000`, the answer is `0`
* Otherwise, every integer from `1000` through `n` has exactly one comma

So the answer is simply the count of integers in `[1000, n]`:

```text
n - 999
```

## Algorithm

1. If `n < 1000`, return `0`.
2. Otherwise return `n - 999`.

## Why This Works

### Why only one comma appears in this range

For numbers up to `100,000`, standard formatting uses:

* `999` and below: no comma
* `1,000` through `99,999`: one comma
* `100,000`: one comma

Since `n <= 10^5`, we never reach values such as `1,000,000` that would require two commas.

### Why counting integers is enough

Each qualifying integer contributes exactly one comma, so total commas equal the number of integers from `1000` to `n` inclusive:

```text
n - 1000 + 1 = n - 999
```

## Complexity Analysis

### Time Complexity

```text
O(1)
```

The answer is computed with a constant number of operations.

### Space Complexity

```text
O(1)
```

No extra data structures are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Return early when `n < 1000`
* Use `n - 999` for the inclusive count from `1000` to `n`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Minimum input
* Last value without commas
* First value with one comma
* Upper bound of the four-digit range
* Five-digit input
* Maximum constraint value

## Edge Cases

Important cases to consider:

* `n = 1`
* `n = 999`
* `n = 1000`
* `n = 9999`
* `n = 100000`

## Notes

* A simulation that formats every integer also works, but the formula is faster and simpler.
* The key observation is that the constraint cap keeps every formatted number at one comma only.
