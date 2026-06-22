---
id: 2287
title: "Rearrange Characters to Make Target String"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/rearrange-characters-to-make-target-string/"
contest: "Weekly Contest 295"
status: "Solved"
language: "Go"
topics:
  - "Hash Table"
  - "String"
  - "Counting"
go_concepts:
  - "Fixed-size arrays for letter counts"
  - "String iteration by byte index"
  - "Integer division for frequency limits"
  - "Taking the minimum across counts"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - hash-table
  - string
  - counting
  - weekly-contest-295
---

# 2287. Rearrange Characters to Make Target String

## Problem Link

LeetCode: `https://leetcode.com/problems/rearrange-characters-to-make-target-string/`

## Difficulty

Easy

## Problem Topics

* Hash Table
* String
* Counting

## What to Know Before Solving

General concepts:

* Count how many times each letter appears in `s` and in `target`
* Each copy of `target` consumes its letter frequencies from `s`
* The answer is limited by the scarcest required letter
* Letters in `s` that are not needed by `target` can be ignored

Go concepts:

* Using `[26]int` as a fixed-size frequency table
* Mapping a lowercase letter to an index with `ch - 'a'`
* Integer division to compute how many copies one letter supports
* Comparing values to find a minimum

## Problem Description

You are given two 0-indexed strings `s` and `target`.

You can take some letters from `s` and rearrange them to form new strings.

Return the maximum number of copies of `target` that can be formed by taking letters from `s` and rearranging them.

## Function Signature

Expected LeetCode function signature:

```go
func rearrangeCharacters(s string, target string) int {

}
```

## Examples

### Example 1

Input:

```text
s = "ilovecodingonleetcode", target = "code"
```

Output:

```text
2
```

Explanation:

```text
Two copies of "code" can be formed from the letters in s.
```

### Example 2

Input:

```text
s = "abcba", target = "abc"
```

Output:

```text
1
```

Explanation:

```text
Only one copy of "abc" can be formed because the single "c" cannot be reused.
```

### Example 3

Input:

```text
s = "abbaccaddaeea", target = "aaaaa"
```

Output:

```text
1
```

Explanation:

```text
Only one copy of "aaaaa" can be formed.
```

## Constraints

```text
1 <= s.length <= 100
1 <= target.length <= 10
s and target consist of lowercase English letters
```

## Approach

Count the frequency of every lowercase letter in `s` and in `target`.

For each letter that appears in `target`, compute how many copies of `target` that letter alone can support:

```text
sourceCount[letter] / targetCount[letter]
```

The answer is the minimum of those values.

A simple upper bound is `len(s) / len(target)`, which can be used as the initial result before checking each letter.

## Algorithm

1. Build `sourceCount` from `s`.
2. Build `targetCount` from `target`.
3. Set `result = len(s) / len(target)`.
4. For each letter index from 0 to 25:
   1. Skip letters that do not appear in `target`.
   2. Compute `copies = sourceCount[i] / targetCount[i]`.
   3. Update `result` if `copies` is smaller.
5. Return `result`.

## Why This Works

Each copy of `target` needs a fixed multiset of letters.

If `s` contains `sourceCount[c]` copies of a letter `c`, and each `target` needs `targetCount[c]` copies of that letter, then letter `c` alone can support at most:

```text
sourceCount[c] / targetCount[c]
```

copies of `target`.

All letters must be available at the same time, so the final answer is the minimum across every letter required by `target`.

## Complexity Analysis

Let `n` be the length of `s` and let `m` be the length of `target`.

### Time Complexity

```text
O(n + m)
```

Both strings are scanned once, and the final loop over 26 letters is constant.

### Space Complexity

```text
O(1)
```

Two fixed-size arrays of length 26 are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `[26]int` is enough because the input uses lowercase English letters only
* Letters with zero demand in `target` are skipped
* `len(s) / len(target)` gives a safe initial upper bound

```go
package rearrangecharacterstomaketargetstring

func rearrangeCharacters(s string, target string) int {
	var sourceCount [26]int
	var targetCount [26]int

	for i := 0; i < len(s); i++ {
		sourceCount[s[i]-'a']++
	}

	for i := 0; i < len(target); i++ {
		targetCount[target[i]-'a']++
	}

	result := len(s) / len(target)

	for i := 0; i < 26; i++ {
		if targetCount[i] == 0 {
			continue
		}

		copies := sourceCount[i] / targetCount[i]

		if copies < result {
			result = copies
		}
	}

	return result
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All LeetCode examples
* Equivalent `balloon` cases from problem 1189
* No valid copy
* Exact single match
* Single-character `target`
* Missing required letter

## Edge Cases

Important cases to consider:

* `target` cannot be formed at all
* `s` and `target` are identical
* `target` contains repeated letters
* `s` has extra unused letters
* One letter in `target` is the bottleneck

## Notes

* This problem is the same as LeetCode 1189: Maximum Number of Balloons when `target = "balloon"`.
* Problem 1189 uses direct counters for the fixed word `balloon`.
* This generalized version uses frequency arrays and works for any `target`.
