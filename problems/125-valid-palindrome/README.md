---
id: 125
title: "Valid Palindrome"
difficulty: "Easy"
level: "Beginner"
platform: "LeetCode"
link: "https://leetcode.com/problems/valid-palindrome/"
status: "Solved"
language: "Go"
topics:
  - "Two Pointers"
  - "String"
go_concepts:
  - "String indexing with bytes"
  - "Two pointers from both ends"
  - "ASCII letter and digit checks"
  - "Case normalization without unicode package"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - two-pointers
  - string
---

# 125. Valid Palindrome

## Problem Link

LeetCode: `https://leetcode.com/problems/valid-palindrome/`

## Difficulty

Easy

## Problem Topics

* Two Pointers
* String

## What to Know Before Solving

General concepts:

* What counts as alphanumeric (letters `A`–`Z`, `a`–`z`, digits `0`–`9`)
* Case-insensitive comparison after normalizing to lowercase
* Palindrome definition on the cleaned sequence

Go concepts:

* Treating `string` as a read-only byte sequence for ASCII input
* Moving `left` and `right` indices with `for` loops
* Small helper functions for readability (`isAlphaNum`, `toLower`)

## Problem Description

A phrase is a palindrome if, after converting all uppercase letters to lowercase and removing all non-alphanumeric characters, it reads the same forward and backward.

Given a string `s`, return `true` if it is a palindrome, and `false` otherwise.

## Function Signature

Expected LeetCode function signature:

```go
func isPalindrome(s string) bool {

}
```

## Examples

### Example 1

Input:

```text
s = "A man, a plan, a canal: Panama"
```

Output:

```text
true
```

Explanation:

```text
After cleaning: "amanaplanacanalpanama", which is a palindrome.
```

### Example 2

Input:

```text
s = "race a car"
```

Output:

```text
false
```

Explanation:

```text
After cleaning: "raceacar", which is not a palindrome.
```

### Example 3

Input:

```text
s = " "
```

Output:

```text
true
```

Explanation:

```text
After removing non-alphanumeric characters, the string is empty.
An empty string is a palindrome.
```

## Constraints

```text
1 <= s.length <= 2 * 10^5
s consists only of printable ASCII characters.
```

## Approach

Use two pointers starting at the beginning and end of `s`. Skip characters that are not alphanumeric on each side. Compare the two remaining characters using the same case (lowercase). Move inward and repeat.

If any comparison fails, return `false`. If the pointers cross without failure, return `true`.

## Algorithm

1. Set `left` to `0` and `right` to `len(s) - 1`.
2. While `left < right`:

    1. Advance `left` while it points to a non-alphanumeric byte and `left < right`.
    2. Move `right` backward while it points to a non-alphanumeric byte and `left < right`.
    3. If `toLower(s[left]) != toLower(s[right])`, return `false`.
    4. Increment `left` and decrement `right`.
3. Return `true`.

## Why This Works

Cleaning the string produces a subsequence of alphanumeric characters in their original order. The two-pointer walk compares the first remaining character with the last remaining character, then the next pair inward, which is exactly the palindrome check on that subsequence.

Skipping non-alphanumeric characters never removes a character that should participate in the comparison, so the result matches the cleaned string palindrome definition.

## Complexity Analysis

Let `n` be the length of `s`.

### Time Complexity

```text
O(n)
```

Each index is visited a constant number of times because `left` only increases and `right` only decreases.

### Space Complexity

```text
O(1)
```

Only indices and a few local variables are used; no extra copy of the string is required.

## Code

The Go solution is available in:

```text
solution.go
```

Two-pointer solution with ASCII helpers:

```go
package validpalindrome

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

## Edge Cases

Important cases to consider:

* Only spaces or punctuation (cleaned string is empty)
* Single effective character after cleaning
* Letters with different cases that should match
* Digits mixed with letters
* Minimum and maximum length strings allowed by constraints
