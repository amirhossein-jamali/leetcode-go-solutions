---
id: 3700
title: "Number of ZigZag Arrays II"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/number-of-zigzag-arrays-ii/"
contest: "Weekly Contest 469"
status: "Solved"
language: "Go"
topics:
  - "Math"
  - "Dynamic Programming"
go_concepts:
  - "int64 modular arithmetic"
  - "Structured matrix-vector multiplication in O(q)"
  - "Linear recurrence from Cayley-Hamilton theorem"
  - "Polynomial binary exponentiation"
  - "Factorials and inverse factorials for binomial coefficients"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - math
  - dynamic-programming
  - linear-algebra
  - matrix-exponentiation
  - cayley-hamilton
  - weekly-contest-469
---

# 3700. Number of ZigZag Arrays II

## Problem Link

LeetCode: `https://leetcode.com/problems/number-of-zigzag-arrays-ii/`

## Difficulty

Hard

## Problem Topics

* Math
* Dynamic Programming

## What to Know Before Solving

General concepts:

* ZigZag arrays alternate between increasing and decreasing adjacent comparisons
* Part I (`3699`) can simulate length directly because `n <= 2000`
* Part II allows `n <= 10^9`, so the answer must come from recurrence or exponentiation
* Only relative order matters, so values can be normalized to `0..q`
* The transition matrix `M[i,j] = min(i,j)` has special structure
* Cayley-Hamilton converts matrix powers into a linear recurrence of order `q`
* Polynomial binary exponentiation evaluates the `t`-th term in `O(q^2 log n)` time

Go concepts:

* Using `int64` for all modular arithmetic
* Building initial sequence terms with fast `M * vector` multiplication
* Computing recurrence coefficients with factorials and inverse factorials
* Reducing polynomials modulo the characteristic polynomial
* Handling the special case `q = 1` separately
* Writing table-driven tests with the `testing` package

## Problem Description

You are given three integers `n`, `l`, and `r`.

A ZigZag array of length `n` is defined as follows:

* Each element lies in the range `[l, r]`.
* No two adjacent elements are equal.
* No three consecutive elements form a strictly increasing or strictly decreasing sequence.

Return the total number of valid ZigZag arrays.

Since the answer may be large, return it modulo `10^9 + 7`.

This is the large-`n` version of [3699. Number of ZigZag Arrays I](./../3699-number-of-zigzag-arrays-i/).

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
3 <= n <= 10^9
1 <= l < r <= 75
```

## Approach

### Intuition

Adjacent comparisons must alternate between `<` and `>`. Otherwise three consecutive elements would be strictly monotone.

Every valid array is either:

```text
a1 < a2 > a3 < a4 > ...
```

or:

```text
a1 > a2 < a3 > a4 < ...
```

These two families are symmetric, so we count only arrays beginning with `<` and multiply by `2`.

### Normalize the value range

Let:

```text
m = r - l + 1
q = r - l
```

Only comparisons matter, so map `[l, r]` to `0, 1, ..., q`.

### Two-step transition matrix

Instead of adding one comparison at a time, combine two steps:

```text
peak -> valley -> peak
```

If the current peak is `i` and the next peak is `j`, the valley `x` must satisfy:

```text
x < min(i, j)
```

So there are exactly `min(i, j)` choices for `x`.

Define:

```text
M[j, i] = min(i, j)
```

Multiplying `M` by a state vector extends the array length by `2`.

### Initial vector

For increasing pairs `a1 < a2` ending at value `i`, there are `i` choices for the first value.

So:

```text
v = [1, 2, 3, ..., q]
```

### Even and odd lengths

Let:

```text
t = (n - 2) / 2   // integer division
```

For even `n`:

```text
answer/2 = 1^T * M^t * v
```

For odd `n`, integer division gives `t = (n - 3) / 2` and one final decreasing step is needed:

```text
answer/2 = v^T * M^t * v
```

The same exponent expression works for both parities.

### Fast matrix-vector multiplication

Although `M` is `q x q`, multiplication by a vector takes `O(q)` time because:

```text
(Mx)_i = sum_{j=1..i} j*x_j + i * sum_{j=i+1..q} x_j
```

This uses weighted prefix sums and total sums.

### Cayley-Hamilton and linear recurrence

Direct matrix exponentiation would cost `O(q^3 log n)`.

The matrix `M[i,j] = min(i,j)` satisfies a degree-`q` characteristic polynomial. By Cayley-Hamilton:

```text
M^q = c0*M^(q-1) + c1*M^(q-2) + ... + c_{q-1}*I
```

Therefore every sequence:

```text
S_t = a^T * M^t * b
```

obeys a linear recurrence of order `q`.

The recurrence coefficients are:

```text
c_j = (-1)^j * C(q + 1 + j, q - 1 - j)
```

We:

1. Build the first `q` terms of the required sequence in `O(q^2)` time.
2. Compute `x^t mod P(x)` with polynomial binary exponentiation in `O(q^2 log n)` time.
3. Evaluate the answer as a dot product with the initial terms.

### Special case `q = 1`

When the interval has only two values, every valid array is forced after the first choice. Exactly `2` arrays exist for every allowed length.

## Algorithm

1. Set `q = r - l`.
2. If `q == 1`, return `2`.
3. Compute `exponent = (n - 2) / 2` and detect odd length with `n & 1`.
4. Build the first `q` sequence terms:
   * Start with `current = [1, 2, ..., q]`.
   * For even length, each term is the sum of `current`.
   * For odd length, each term is `sum(i * current[i-1])`.
   * Update `current <- M * current` using fast multiplication.
5. If `exponent < q`, return `2 * initial[exponent]`.
6. Build recurrence coefficients with binomial formulas modulo `mod`.
7. Evaluate the `exponent`-th term with polynomial binary exponentiation.
8. Return `2 * oneDirection mod mod`.

## Why This Works

Alternating comparisons are exactly the ZigZag constraint.

The bijection `x -> l + r - x` shows that increasing-first and decreasing-first arrays are equinumerous.

The matrix `M[j,i] = min(i,j)` counts all valid two-step extensions from peak `i` to peak `j`.

The scalar sequences used for even and odd lengths both have the form `a^T M^t b`, so Cayley-Hamilton gives the same recurrence for both.

Generating the first `q` terms provides enough initial data for that recurrence.

Polynomial exponentiation computes `M^t` implicitly without storing dense matrices.

Multiplying by `2` at the end accounts for both starting directions.

## Complexity Analysis

Let:

```text
q = r - l
```

### Time Complexity

```text
O(q^2 log n)
```

Building the first `q` terms takes `O(q^2)`.

Polynomial binary exponentiation takes `O(q^2 log n)`.

### Space Complexity

```text
O(q)
```

Only a constant number of length-`q` arrays are stored.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `exponent = (n - 2) / 2` works for both even and odd `n`.
* `multiplyMinMatrixVector` implements `O(q)` matrix-vector multiplication.
* `buildRecurrence` uses factorials and inverse factorials for binomial coefficients.
* `nthLinearRecurrence` performs polynomial binary exponentiation modulo the characteristic polynomial.
* `q == 1` is handled separately because only two values exist.

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* The two-value special case with large `n`
* Cross-check with part I for `n = 5`
* Even length and very large `n`

## Edge Cases

Important cases to consider:

* `q = 1` with huge `n`
* Small `n` where the answer can be read directly from the initial sequence
* Even and odd `n`
* Maximum `q = 74` when `l = 1` and `r = 75`
* Large `n` up to `10^9`

## Notes

* This is the optimized follow-up to problem `3699`.
* The key optimization is Cayley-Hamilton plus polynomial exponentiation instead of dense matrix exponentiation.
* The matrix structure `M[i,j] = min(i,j)` is what makes both fast multiplication and the closed-form characteristic polynomial possible.
* Related ideas: linear recurrences, matrix exponentiation, combinatorics, prefix sums.

## Related Problem

* [3699. Number of ZigZag Arrays I](./../3699-number-of-zigzag-arrays-i/) — same rules, but `n <= 2000`, solved with direct DP simulation.
