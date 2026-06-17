---
id: 3614
title: "Process String with Special Operations II"
difficulty: "Hard"
level: "Advanced"
platform: "LeetCode"
link: "https://leetcode.com/problems/process-string-with-special-operations-ii/"
contest: "Weekly Contest 458"
status: "Solved"
language: "Go"
topics:
  - "String"
  - "Simulation"
go_concepts:
  - "int64 for large indices and lengths"
  - "Forward pass to track result length"
  - "Backward pass to locate the kth character"
  - "Modulo arithmetic to fold duplicate indices"
  - "Switch with boolean cases"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - string
  - simulation
  - weekly-contest-458
---

# 3614. Process String with Special Operations II

## Problem Link

LeetCode: `https://leetcode.com/problems/process-string-with-special-operations-ii/`

## Difficulty

Hard

## Problem Topics

* String
* Simulation

## What to Know Before Solving

General concepts:

* The same special operations as problem 3612, but the final string can be astronomically large
* Why building the full result string is impossible when length can reach `10^15`
* How to track only the length in a forward pass
* How to walk backward through operations to find one character at index `k`
* How duplicate and reverse operations transform indices

Go concepts:

* Using `int64` for `k` and length because constraints exceed 32-bit range
* Forward simulation of length only, without materializing the string
* Backward traversal with index remapping
* Using `k %= half` when undoing a duplicate operation
* Using `k = length - 1 - k` when undoing a reverse operation
* Returning `byte` for a single character result
* Writing table-driven tests with the `testing` package

## Problem Description

You are given a string `s` consisting of lowercase English letters and the special characters: `*`, `#`, and `%`.

You are also given an integer `k`.

Build a new string `result` by processing `s` according to the following rules from left to right:

* If the character is a lowercase English letter, append it to `result`.
* A `*` removes the last character from `result`, if it exists.
* A `#` duplicates the current `result` and appends it to itself.
* A `%` reverses the current `result`.

Return the `k`th character of the final string `result`. If `k` is out of the bounds of `result`, return `'.'`.

## Function Signature

Expected LeetCode function signature:

```go
func processStr(s string, k int64) byte {

}
```

## Examples

### Example 1

Input:

```text
s = "a#b%*", k = 1
```

Output:

```text
"a"
```

Explanation:

```text
i  s[i]  Operation                 Current result
0  'a'   Append 'a'                "a"
1  '#'   Duplicate result          "aa"
2  'b'   Append 'b'                "aab"
3  '%'   Reverse result            "baa"
4  '*'   Remove the last character "ba"
```

The final result is `"ba"`. The character at index `k = 1` is `'a'`.

### Example 2

Input:

```text
s = "cd%#*#", k = 3
```

Output:

```text
"d"
```

Explanation:

```text
i  s[i]  Operation                 Current result
0  'c'   Append 'c'                "c"
1  'd'   Append 'd'                "cd"
2  '%'   Reverse result            "dc"
3  '#'   Duplicate result          "dcdc"
4  '*'   Remove the last character "dcd"
5  '#'   Duplicate result          "dcddcd"
```

The final result is `"dcddcd"`. The character at index `k = 3` is `'d'`.

### Example 3

Input:

```text
s = "z*#", k = 0
```

Output:

```text
"."
```

Explanation:

```text
i  s[i]  Operation        Current result
0  'z'   Append 'z'       "z"
1  '*'   Remove last char ""
2  '#'   Duplicate result ""
```

The final result is `""`. Since index `k = 0` is out of bounds, the output is `'.'`.

## Constraints

```text
1 <= s.length <= 10^5
s consists of only lowercase English letters and special characters *, #, and %.
0 <= k <= 10^15
The length of result after processing s will not exceed 10^15.
```

## Approach

This problem extends 3612 with one critical difference: the final string may be far too large to build.

Instead of constructing `result`, we use two passes:

1. **Forward pass:** simulate only the final length after each operation.
2. **Backward pass:** start from the final length and walk `s` from right to left, remapping `k` until it points to a source letter in the original input.

If `k >= length`, return `'.'` immediately.

## Algorithm

### Pass 1: compute final length

1. Set `length = 0`.
2. Scan `s` from left to right:
   * Letter: `length++`
   * `*`: if `length > 0`, `length--`
   * `#`: `length *= 2`
   * `%`: no length change
3. If `k >= length`, return `'.'`.

### Pass 2: locate the kth character

1. Scan `s` from right to left.
2. For each character, undo the operation on `(k, length)`:
   * Letter: if `k == length-1`, return that letter; otherwise `length--`
   * `*`: undo delete by `length++`
   * `#`: let `half = length / 2`, set `k %= half`, then `length = half`
   * `%`: reverse index with `k = length - 1 - k`
3. If no letter is found, return `'.'`.

## Why This Works

The forward pass tracks exactly how long `result` becomes, without storing the characters.

When walking backward:

* A letter at the end of the current virtual string is identified when `k == length-1`.
* Undoing `*` restores a deleted trailing character, so length increases by one.
* Undoing `#` means the second half is a copy of the first half, so any index in the second half folds into the first half with `k %= half`.
* Undoing `%` mirrors indices, so `k` becomes `length - 1 - k`.

Each backward step restores the state before that operation was applied. Eventually `k` lands on the original source letter that produced the answer.

## Complexity Analysis

Let `n` be the length of `s`.

### Time Complexity

```text
O(n)
```

Both the forward and backward passes scan `s` once.

### Space Complexity

```text
O(1)
```

Only a few integer variables are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `int64` is required because `k` and the final length can be as large as `10^15`.
* The backward pass never builds the full result string.
* `k %= half` is the key step for undoing duplicate operations efficiently.

```go
package processstringwithspecialoperationsii

func processStr(s string, k int64) byte {
	var length int64

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch {
		case ch >= 'a' && ch <= 'z':
			length++
		case ch == '*':
			if length > 0 {
				length--
			}
		case ch == '#':
			length *= 2
		case ch == '%':
		}
	}

	if k >= length {
		return '.'
	}

	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]

		switch {
		case ch >= 'a' && ch <= 'z':
			if k == length-1 {
				return ch
			}
			length--

		case ch == '*':
			length++

		case ch == '#':
			half := length / 2
			k %= half
			length = half

		case ch == '%':
			k = length - 1 - k
		}
	}

	return '.'
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* First and out-of-bounds indices
* Single-letter input
* Empty result after delete
* Duplicate and reverse index remapping

## Edge Cases

Important cases to consider:

* `k` equal to `length - 1`
* `k` equal to `length` or greater
* Empty final result
* `*` on an empty result
* `#` after operations that leave a large virtual string
* `%` followed by duplicate, which changes how indices fold
* Very large `k` values that still fit in `int64`

## Notes

* This is the optimized follow-up to 3612. The brute-force simulation from part I is too slow and too memory-heavy here.
* The backward pass is a standard technique for "find kth character without building the string" problems.
* Pay special attention to undoing `#` with modulo and undoing `%` with index reflection.
