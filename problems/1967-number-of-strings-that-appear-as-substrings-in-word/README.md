---
id: 1967
title: "Number of Strings That Appear as Substrings in Word"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/"
contest: "Weekly Contest 254"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "String"
go_concepts:
  - "Iterating over a slice of strings"
  - "strings.Contains for substring checks"
  - "Simple counting with an integer variable"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - string
  - weekly-contest-254
---

# 1967. Number of Strings That Appear as Substrings in Word

## Problem Link

LeetCode: `https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/`

## Difficulty

Easy

## Problem Topics

* Array
* String

## What to Know Before Solving

General concepts:

* A substring is a contiguous sequence of characters inside a string
* Each pattern must be checked independently against `word`
* Duplicate patterns are counted separately if they appear in the input array
* `strings.Contains` is enough for this problem size

Go concepts:

* Iterating over `[]string` with `for _, pattern := range patterns`
* Using `strings.Contains(word, pattern)` from the standard library
* Incrementing a counter when a match is found
* Writing table-driven tests with the `testing` package

## Problem Description

Given an array of strings `patterns` and a string `word`, return the number of strings in `patterns` that exist as a substring in `word`.

A substring is a contiguous sequence of characters within a string.

## Function Signature

Expected LeetCode function signature:

```go
func numOfStrings(patterns []string, word string) int {

}
```

## Examples

### Example 1

Input:

```text
patterns = ["a","abc","bc","d"], word = "abc"
```

Output:

```text
3
```

Explanation:

```text
- "a" appears as a substring in "abc".
- "abc" appears as a substring in "abc".
- "bc" appears as a substring in "abc".
- "d" does not appear as a substring in "abc".

3 of the strings in patterns appear as a substring in word.
```

### Example 2

Input:

```text
patterns = ["a","b","c"], word = "aaaaabbbbb"
```

Output:

```text
2
```

Explanation:

```text
- "a" appears as a substring in "aaaaabbbbb".
- "b" appears as a substring in "aaaaabbbbb".
- "c" does not appear as a substring in "aaaaabbbbb".

2 of the strings in patterns appear as a substring in word.
```

### Example 3

Input:

```text
patterns = ["a","a","a"], word = "ab"
```

Output:

```text
3
```

Explanation:

```text
Each of the patterns appears as a substring in word "ab".
```

## Constraints

```text
1 <= patterns.length <= 100
1 <= patterns[i].length <= 100
1 <= word.length <= 100
patterns[i] and word consist of lowercase English letters.
```

## Approach

Check each pattern one by one.

If `pattern` appears inside `word` as a contiguous substring, increment the answer.

Because the constraints are small, a direct substring check for every pattern is sufficient.

## Algorithm

1. Initialize `count = 0`.
2. Iterate over every `pattern` in `patterns`.
3. If `strings.Contains(word, pattern)` is true, increment `count`.
4. Return `count`.

## Why This Works

The problem asks for the number of patterns that appear as substrings in `word`.

`strings.Contains` returns true exactly when `pattern` occurs contiguously inside `word`.

Checking each pattern independently and counting matches therefore gives the required answer.

Duplicate patterns in the input array are evaluated separately, so each one can contribute to the count.

## Complexity Analysis

Let `p` be the number of patterns, and let `L` be the maximum length among `patterns[i]` and `word`.

### Time Complexity

```text
O(p * L^2)
```

In the worst case, each `strings.Contains` scan is linear in `word` length, and pattern comparison adds another factor. With the given constraints, this is easily fast enough.

### Space Complexity

```text
O(1)
```

Only a counter and loop variables are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `strings.Contains` handles empty-pattern edge cases correctly in Go
* The solution counts every pattern entry separately, including duplicates

```go
package numberofstringsthatappearsassubstringsinword

import "strings"

func numOfStrings(patterns []string, word string) int {
	count := 0

	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}

	return count
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* No matching patterns
* Exact full-word match

## Edge Cases

Important cases to consider:

* Duplicate patterns in the input array
* Pattern equal to the entire word
* Single-character patterns
* Patterns that do not appear at all
* Pattern shorter than `word` but matching a prefix or suffix

## Notes

* This is a straightforward string-checking problem.
* No hash map or advanced string algorithm is needed because the constraints are small.
* If the constraints were much larger, more advanced substring search methods might be considered.
