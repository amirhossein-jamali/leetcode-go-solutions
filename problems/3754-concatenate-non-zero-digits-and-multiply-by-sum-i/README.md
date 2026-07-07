---
id: 3754
title: "Concatenate Non-Zero Digits and Multiply by Sum I"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/"
contest: "Weekly Contest 477"
status: "Solved"
language: "Go"
topics:
  - "Math"
go_concepts:
  - "Digit extraction with modulo and division"
  - "Building a number from least significant digit upward"
  - "int64 for intermediate and return values"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - math
  - digit-manipulation
  - weekly-contest-477
---

# 3754. Concatenate Non-Zero Digits and Multiply by Sum I

## Problem Link

LeetCode: `https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/`

## Difficulty

Easy

## Problem Topics

* Math

## What to Know Before Solving

General concepts:

* Extracting digits from an integer with `% 10` and `/= 10`
* Skipping zero digits while preserving original order
* Building a concatenated number from selected digits
* Summing digits and multiplying two derived values

Go concepts:

* Using `int64` for the result and intermediate values
* A single loop with `n % 10` and `n /= 10`
* Tracking a decimal place multiplier while constructing `x`
* Writing table-driven tests with the `testing` package

## Problem Description

You are given an integer `n`.

Form a new integer `x` by concatenating all the non-zero digits of `n` in their original order. If there are no non-zero digits, `x = 0`.

Let `sum` be the sum of digits in `x`.

Return `x * sum`.

## Function Signature

Expected LeetCode function signature:

```go
func sumAndMultiply(n int) int64 {

}
```

## Examples

### Example 1

Input:

```text
n = 10203004
```

Output:

```text
12340
```

Explanation:

```text
The non-zero digits are 1, 2, 3, and 4. Thus, x = 1234.
The sum of digits is sum = 1 + 2 + 3 + 4 = 10.
Therefore, the answer is x * sum = 1234 * 10 = 12340.
```

### Example 2

Input:

```text
n = 1000
```

Output:

```text
1
```

Explanation:

```text
The non-zero digit is 1, so x = 1 and sum = 1.
Therefore, the answer is x * sum = 1 * 1 = 1.
```

## Constraints

```text
0 <= n <= 10^9
```

## Approach

Traverse `n` from least significant digit to most significant digit.

For each non-zero digit:

* Add it to `sum`
* Append it to `x` by multiplying the digit by `place` and adding to `x`
* Multiply `place` by `10`

Processing from right to left with increasing place values preserves the original left-to-right digit order in `x`.

Return `x * sum`.

## Algorithm

1. Initialize `x = 0`, `sum = 0`, and `place = 1`.
2. While `n > 0`:
   * Read `digit = n % 10`.
   * If `digit != 0`, update `x`, `sum`, and `place`.
   * Set `n = n / 10`.
3. Return `x * sum`.

## Why This Works

The digits of `n` are visited from least significant to most significant.

When a non-zero digit is found, multiplying it by the current `place` and adding it to `x` appends that digit to the constructed number in the correct position.

Because `place` starts at `1` and grows by factors of `10`, the first non-zero digit encountered from the right becomes the ones place, the next becomes the tens place, and so on. This reproduces the original left-to-right order of non-zero digits.

`sum` accumulates exactly the digits that appear in `x`, so the final multiplication uses the required values.

If `n = 0` or every digit is zero, the loop never updates `x` or `sum`, so the answer is `0`.

## Complexity Analysis

Let `d` be the number of digits in `n`.

### Time Complexity

```text
O(d)
```

Each digit is processed once.

### Space Complexity

```text
O(1)
```

Only a constant number of variables are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `int64` avoids overflow for intermediate values within the constraints
* `place` builds `x` without string conversion
* Zero digits are skipped without affecting order

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* `n = 0`
* Single-digit input
* Input with no zeros
* Input with many zeros
* Maximum constraint value

## Edge Cases

Important cases to consider:

* `n = 0`
* `n` with only one non-zero digit
* `n` with no zero digits
* `n` with zeros between non-zero digits
* `n = 10^9`

## Notes

* A string-based approach would also work but allocates extra memory.
* The digit loop is the natural simulation of the problem statement.
* Part II of this problem family may require larger values or different rules.
