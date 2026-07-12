---
id: 1331
title: "Rank Transform of an Array"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/rank-transform-of-an-array/"
contest: "Biweekly Contest 18"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
  - "Sorting"
go_concepts:
  - "LSD radix sort on indices"
  - "Ping-pong buffers"
  - "Counting sort buckets"
  - "Signed-to-unsigned key transform"
  - "Stable index sorting"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - sorting
  - radix-sort
  - counting-sort
  - biweekly-contest-18
---

# 1331. Rank Transform of an Array

## Problem Link

LeetCode: `https://leetcode.com/problems/rank-transform-of-an-array/`

## Difficulty

Easy

## Problem Topics

* Array
* Hash Table
* Sorting

## What to Know Before Solving

General concepts:

* Rank starts at `1` for the smallest value
* Equal values must receive the same rank
* Larger values must receive larger ranks
* The task is equivalent to compressing sorted distinct values into consecutive ranks

Go concepts:

* Sorting indices instead of copying values
* LSD radix sort with 8-bit passes
* Ping-pong buffers in one allocation
* Converting signed integers to sortable unsigned keys
* Table-driven tests with a map-based reference

## Problem Description

Given an integer array `arr`, replace each element with its rank.

Rank rules:

1. Rank starts from `1`.
2. Larger values get larger ranks.
3. Equal values get the same rank.
4. Ranks must be as small as possible.

Return the transformed array.

## Function Signature

Expected LeetCode function signature:

```go
func arrayRankTransform(arr []int) []int {

}
```

## Examples

### Example 1

Input:

```text
arr = [40,10,20,30]
```

Output:

```text
[4,1,2,3]
```

Explanation:

```text
10 is smallest, so rank 1.
20 is second smallest, so rank 2.
30 is third smallest, so rank 3.
40 is largest, so rank 4.
```

### Example 2

Input:

```text
arr = [100,100,100]
```

Output:

```text
[1,1,1]
```

### Example 3

Input:

```text
arr = [37,12,28,9,100,56,80,5,12]
```

Output:

```text
[5,3,4,2,8,6,7,1,3]
```

## Constraints

```text
0 <= arr.length <= 10^5
-10^9 <= arr[i] <= 10^9
```

## Approach Comparison

| Method | Time | Extra memory | Notes |
| --- | ---: | ---: | --- |
| Copy + sort + `map` | `O(n log n)` | `O(n)` | Simplest and easiest to maintain |
| Sort `(value, index)` pairs | `O(n log n)` | `O(n)` | No map, but still needs a pair slice |
| In-place arithmetic packing | `O(n log n)` | `O(1)` heap | Low memory, but harder to read |
| LSD radix sort on indices | `O(n)` | `O(n)` | Best asymptotic time for fixed-width integers |

The chosen solution uses **LSD radix sort on indices** with one `O(n)` auxiliary buffer.

## Approach

Instead of copying values or building a rank map:

1. Keep the original values in `arr`.
2. Sort the **indices** `0..n-1` by the corresponding values.
3. Walk the sorted indices once and assign ranks.
4. Write each rank back to `arr[index]`.

Sorting is done with stable LSD radix sort in four 8-bit passes over 32-bit keys.

## Signed Integer Handling

Radix sort expects unsigned ordering. For signed `int` values, each key is transformed with:

```text
key = uint32(int32(value)) ^ (1 << 31)
```

This flips the sign bit so that:

```text
... < -2 < -1 < 0 < 1 < 2 < ...
```

matches unsigned numeric order. The transform is reversible and preserves duplicate detection.

## Algorithm

1. If `n == 0`, return `arr`.
2. Allocate one buffer of size `2n` and split it into `src` and `dst`.
3. Initialize `src` with indices `0, 1, ..., n-1`.
4. For `shift = 0, 8, 16, 24`:
   * Count bucket frequencies for the current byte of each sorted key.
   * Convert counts to starting positions.
   * Scatter indices from `src` into `dst` by bucket.
   * Swap `src` and `dst`.
5. After four passes, `src` contains indices sorted by value.
6. Traverse `src`:
   * Increase rank only when the transformed key changes.
   * Write `arr[index] = rank`.
7. Return `arr`.

## Walkthrough

Input:

```text
arr = [40, 10, 20, 30]
```

Initial index array:

```text
src = [0, 1, 2, 3]
```

After radix sorting indices by value, the order becomes:

```text
src = [1, 2, 3, 0]
```

Rank assignment:

```text
index 1 -> value 10 -> rank 1
index 2 -> value 20 -> rank 2
index 3 -> value 30 -> rank 3
index 0 -> value 40 -> rank 4
```

Write ranks back:

```text
arr[1] = 1
arr[2] = 2
arr[3] = 3
arr[0] = 4
```

Final output:

```text
[4, 1, 2, 3]
```

## Why This Works

### Sorting indices preserves values

The original array is never reordered. Only permutation indices are sorted, so each value remains available at `arr[index]` while we discover value order.

### Stable radix sort keeps duplicate handling simple

Equal values produce equal transformed keys. After sorting, they appear consecutively in `src`, so one linear scan assigns the same rank to all of them.

### Rank scan is correct

The scan increases rank only when the key changes. Therefore:

* the smallest distinct value gets rank `1`
* each next distinct value gets the next integer rank
* equal values share the previous rank

Writing `arr[index] = rank` restores ranks to original positions.

## Complexity Analysis

Let:

```text
n = len(arr)
```

### Time Complexity

Each radix pass is `O(n)`. With four 8-bit passes over 32-bit keys:

```text
O(4n) = O(n)
```

The final rank scan is also `O(n)`.

Total:

```text
O(n)
```

### Space Complexity

One auxiliary buffer stores two index arrays:

```go
storage := make([]uint32, 2*n)
```

So:

```text
O(n)
```

The counting array has fixed size `256`, so it does not depend on `n`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `storage` is split into `src` and `dst` for ping-pong radix passes
* indices are stored as `uint32`, which is enough because `n <= 10^5`
* `int32(arr[index])` normalizes each value before the sign-bit flip
* only `src` is used after sorting completes
* ranks are written directly into the input slice

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Empty array and single element
* Sorted and reverse-sorted inputs
* Duplicate values
* Negative values
* Constraint-scale values near `-10^9` and `10^9`
* Randomized differential tests against a map-based reference

## Edge Cases

Important cases to consider:

* `len(arr) == 0`
* `len(arr) == 1`
* All elements equal
* Many duplicates mixed with distinct values
* Negative numbers
* Values near `-10^9` and `10^9`
* Already sorted or reverse-sorted input

## Notes

* The standard map-based solution is easier to explain in interviews.
* Radix sort is a better fit here when linear time matters and one `O(n)` buffer is acceptable.
* The sign-bit flip trick avoids separate handling for negative numbers during counting sort.
* This solution mutates the input slice in place when writing ranks.
* In-place arithmetic packing is still useful when auxiliary `O(n)` memory must be avoided, but radix sort is simpler to reason about than cycle permutation.

## Completion Checklist

* YAML frontmatter is filled
* Problem link, difficulty, topics, and contest are filled
* Approach explains radix sort on indices
* Function signature matches LeetCode
* Examples and constraints are included
* Time and space complexity are included
* `solution.go` uses the expected LeetCode function signature
* `solution_test.go` covers examples and edge cases
* `go test ./...` passes
