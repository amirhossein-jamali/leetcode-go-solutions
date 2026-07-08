---
id: 3756
title: "Concatenate Non-Zero Digits and Multiply by Sum II"
difficulty: "Medium"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/"
contest: "Weekly Contest 477"
status: "Solved"
language: "Go"
topics:
  - "Math"
  - "String"
  - "Prefix Sum"
go_concepts:
  - "Prefix arrays for count, sum, and concatenated value"
  - "Modular arithmetic for large integers"
  - "Bit packing in uint64 metadata"
  - "Substring query answering in O(1)"
  - "Table-driven tests with brute-force reference"
tags:
  - leetcode
  - go
  - math
  - string
  - prefix-sum
  - modular-arithmetic
  - weekly-contest-477
---

# 3756. Concatenate Non-Zero Digits and Multiply by Sum II

## Problem Link

LeetCode: `https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/`

## Difficulty

Medium

## Problem Topics

* Math
* String
* Prefix Sum

## What to Know Before Solving

General concepts:

* The same digit-concatenation rule as [3754](./../3754-concatenate-non-zero-digits-and-multiply-by-sum-i/), but on many substring queries
* Prefix preprocessing to answer each query in constant time
* Modular concatenation with powers of 10
* Packing count and digit sum into one metadata array

Go concepts:

* Prefix arrays over a digit string
* `uint64` bit packing for `(count, sum)` pairs
* Modular subtraction with negative correction
* Brute-force reference tests for randomized validation

## Problem Description

You are given a string `s` of length `m` consisting of digits. You are also given a 2D integer array `queries`, where `queries[i] = [li, ri]`.

For each query, extract the substring `s[li..ri]`, then:

1. Form `x` by concatenating all non-zero digits in original order. If none exist, `x = 0`.
2. Let `sum` be the sum of digits in `x`.
3. The answer is `x * sum`, modulo `10^9 + 7`.

Return an array of answers for all queries.

## Function Signature

Expected LeetCode function signature:

```go
func sumAndMultiply(s string, queries [][]int) []int {

}
```

## Examples

### Example 1

Input:

```text
s = "10203004", queries = [[0,7],[1,3],[4,6]]
```

Output:

```text
[12340, 4, 9]
```

### Example 2

Input:

```text
s = "1000", queries = [[0,3],[1,1]]
```

Output:

```text
[1, 0]
```

### Example 3

Input:

```text
s = "9876543210", queries = [[0,9]]
```

Output:

```text
[444444137]
```

Explanation:

```text
x = 987654321
sum = 45
987654321 * 45 = 44444444445
44444444445 mod (10^9 + 7) = 444444137
```

## Constraints

```text
1 <= m == s.length <= 10^5
s consists of digits only.
1 <= queries.length <= 10^5
queries[i] = [li, ri]
0 <= li <= ri < m
```

## Approach

Preprocess the full string once.

For every prefix ending at position `i`, store:

* `cnt`: number of non-zero digits seen so far
* `sum`: sum of those digits
* `val`: concatenation of those digits modulo `MOD`
* `pow10[k]`: `10^k mod MOD` for the `k`th non-zero digit count

For a query `[l, r]`, use prefix differences at `l` and `r + 1`:

* `length = cnt[r+1] - cnt[l]`
* `digitSum = sum[r+1] - sum[l]`
* `x = val[r+1] - val[l] * pow10[length] (mod MOD)`

If `length == 0`, the answer is `0`.

## Algorithm

1. Build `meta`, `val`, and `pow10` in one left-to-right scan.
2. Pack `(cnt, sum)` into `meta[i]` using bit shift `20`.
3. For each query `[l, r]`:
   * Read prefix states at `l` and `r + 1`.
   * If no non-zero digits exist in the range, leave answer `0`.
   * Otherwise compute modular substring concatenation and multiply by digit sum.
4. Return all answers.

## Why This Works

### Prefix concatenation

If `val[i]` is the concatenation of all non-zero digits in `s[0..i-1]` modulo `MOD`, then the substring range corresponds to removing the prefix before `l`.

If the range contains `length` non-zero digits, the value built from `s[l..r]` is:

```text
val[r+1] - val[l] * 10^length
```

modulo `MOD`. The precomputed `pow10[length]` supplies `10^length mod MOD`.

### Prefix sums

`sum` and `cnt` are additive over prefixes, so differences `meta[r+1] - meta[l]` give the count and digit sum inside the query range.

### Bit packing

`cnt` is at most `m <= 10^5`, and digit sum is at most `9 * 10^5`. Both fit in 20 bits, so one `uint64` per prefix stores both values.

## Complexity Analysis

Let:

```text
m = len(s)
q = len(queries)
```

### Time Complexity

```text
O(m + q)
```

Preprocessing is `O(m)`. Each query is answered in `O(1)`.

### Space Complexity

```text
O(m)
```

For `meta`, `val`, and `pow10`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Use half-open prefix indices: query `[l, r]` maps to `[l, r+1)`
* Correct modular subtraction when `x < 0`
* `length == 0` means all zeros in the substring, so answer stays `0`
* `SHIFT = 20` is safe for the given constraints

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Single-character strings
* Queries over all-zero substrings
* Randomized differential tests against brute force

## Edge Cases

Important cases to consider:

* Query on a single `'0'`
* Query on a single non-zero digit
* Substring with zeros between non-zero digits
* Very large concatenated value requiring modulo
* Multiple queries on the same string

## Notes

* This is the query version of [3754. Concatenate Non-Zero Digits and Multiply by Sum I](./../3754-concatenate-non-zero-digits-and-multiply-by-sum-i/).
* Rebuilding each substring naively would be too slow for `10^5` queries.
* Prefix preprocessing trades `O(m)` memory for `O(1)` query time.
