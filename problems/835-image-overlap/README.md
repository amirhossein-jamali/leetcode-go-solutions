---
id: 835
title: "Image Overlap"
difficulty: "Medium"
level: "Senior Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/image-overlap/"
contest: "Weekly Contest 84"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Matrix"
go_concepts:
  - "Bitset row compression with uint32"
  - "math/bits OnesCount32"
  - "Prefix sums for row popcounts"
  - "Translation enumeration with pruning"
  - "Fixed-size stack arrays"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - matrix
  - bit-manipulation
  - weekly-contest-84
---

# 835. Image Overlap

## Problem Link

LeetCode: `https://leetcode.com/problems/image-overlap/`

## Difficulty

Medium

## Problem Topics

* Array
* Matrix

## What to Know Before Solving

General concepts:

* A translation moves every `1` in one image left, right, up, or down without rotation
* Bits that move outside the matrix disappear
* Overlap is the count of positions where both images have a `1`
* With `n <= 30`, each row fits in a 32-bit integer
* There are only `(2n - 1)^2` possible translations, so brute force is feasible with good pruning

Go concepts:

* Compressing each matrix row into `uint32`
* Counting set bits with `bits.OnesCount32`
* Building prefix sums over row popcounts
* Shifting rows left or right to simulate horizontal translation
* Using fixed-size arrays `[30]uint32` and `[31]int`
* Writing table-driven tests with the `testing` package

## Problem Description

You are given two binary square matrices `img1` and `img2` of size `n x n`.

Translate one image by sliding its `1` bits in the four cardinal directions any number of units, place it on top of the other image, and count overlapping `1` positions.

Return the largest possible overlap.

## Function Signature

Expected LeetCode function signature:

```go
func largestOverlap(img1 [][]int, img2 [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
img1 = [[1,1,0],[0,1,0],[0,1,0]]
img2 = [[0,0,0],[0,1,1],[0,0,1]]
```

Output:

```text
3
```

Explanation:

```text
Translate img1 one unit right and one unit down to get overlap 3.
```

### Example 2

Input:

```text
img1 = [[1]]
img2 = [[1]]
```

Output:

```text
1
```

### Example 3

Input:

```text
img1 = [[0]]
img2 = [[0]]
```

Output:

```text
0
```

## Constraints

```text
n == img1.length == img1[i].length
n == img2.length == img2[i].length
1 <= n <= 30
img1[i][j] is either 0 or 1
img2[i][j] is either 0 or 1
```

## Approach

Compress each row of both images into a bitset.

For every horizontal translation of `img1`:

1. Shift all compressed rows left or right.
2. Store the shifted rows and their prefix popcounts.

For every vertical translation:

1. Align the overlapping row ranges of the shifted `img1` and original `img2`.
2. Use prefix sums to prune translations that cannot beat the current answer.
3. Count overlap with bitwise `AND` plus `OnesCount32` on each aligned row.

Stop early when the answer reaches `min(totalOnes(img1), totalOnes(img2))`.

## Algorithm

1. Convert each row of `img1` and `img2` into `uint32` bitmasks.
2. Count total ones in each image; return `0` if either image has no ones.
3. For each horizontal offset `dx`:
   * Build `shifted` rows from `img1`.
   * Build prefix popcounts `ap`.
   * Skip if total ones in `shifted` is not greater than current best.
4. For each vertical offset `dy`:
   * Compute overlapping row interval `[ai, ai+rows)` in `shifted` and `[bi, bi+rows)` in `img2`.
   * Prune using row-range popcounts.
   * Add `popcount(shifted[row] & img2[row])` across overlapping rows.
5. Track and return the maximum overlap.

## Why This Works

### Why row bitsets are enough

Each matrix row has at most 30 bits, so a row can be stored in one `uint32`. Horizontal translation becomes a left or right shift with masking. Vertical translation becomes row alignment only.

### Why prefix popcounts help

If the total number of ones in the overlapping rows of either image is already less than or equal to the current best answer, that translation cannot improve the result. This prunes many `(dx, dy)` pairs cheaply.

### Why early exit is valid

The overlap can never exceed the number of ones in the smaller image. Once the current best reaches that bound, the answer is optimal.

## Complexity Analysis

Let:

```text
n = matrix size
```

There are `O(n^2)` translations.

### Time Complexity

```text
O(n^3)
```

For each of `O(n^2)` translations, we scan up to `n` rows and do constant bit work per row.

With `n <= 30`, this is fast in practice.

### Space Complexity

```text
O(n)
```

Only a few fixed-size arrays of length `O(n)` are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Pack each row with `x |= uint32(value) << j`
* Mask left shifts with `(a[i] << s) & mask`
* Encode horizontal offsets with the `tx` loop pattern
* Encode vertical offsets with the `ty` loop pattern
* Use `bits.OnesCount32` for popcount on row intersections

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Identical small images
* Image with no ones
* Disjoint single ones
* Diagonal patterns with limited overlap

## Edge Cases

Important cases to consider:

* `n = 1`
* Both images all zero
* One image all zero
* Identical images
* Best overlap requires negative horizontal or vertical shift
* Patterns where only one row overlaps

## Notes

* A convolution or hash-based approach also works, but bitsets are simpler for `n <= 30`.
* The translation loops use a compact indexing trick instead of nested loops over raw offsets.
* This problem is a good example of turning matrix translation into bitwise operations.
