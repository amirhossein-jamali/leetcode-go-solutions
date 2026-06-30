---
id: 1358
title: "Number of Substrings Containing All Three Characters"
difficulty: "Medium"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/"
contest: "Biweekly Contest 20"
status: "Solved"
language: "Go"
topics:
  - "Hash Table"
  - "String"
  - "Sliding Window"
go_concepts:
  - "Fixed-size array for character counts"
  - "Sliding window with two pointers"
  - "Byte indexing with s[i]-'a'"
  - "Counting valid substrings by extending right end"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - hash-table
  - string
  - sliding-window
  - biweekly-contest-20
---

# 1358. Number of Substrings Containing All Three Characters

## Problem Link

LeetCode: `https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/`

## Difficulty

Medium

## Problem Topics

* Hash Table
* String
* Sliding Window

## What to Know Before Solving

General concepts:

* A substring must contain at least one `a`, one `b`, and one `c`
* Sliding window can track whether the current window is valid
* Once a window `[left, right]` is valid, every extension to the right is also valid
* Counting substrings ending at `right` avoids double counting

Go concepts:

* Using `[3]int` as a frequency table for three characters
* Expanding `right` and shrinking `left` in a two-pointer loop
* Mapping characters with `s[i] - 'a'`
* Accumulating answer with `result += n - right`
* Writing table-driven tests with the `testing` package

## Problem Description

Given a string `s` consisting only of characters `a`, `b`, and `c`.

Return the number of substrings containing at least one occurrence of all these characters `a`, `b`, and `c`.

## Function Signature

Expected LeetCode function signature:

```go
func numberOfSubstrings(s string) int {

}
```

## Examples

### Example 1

Input:

```text
s = "abcabc"
```

Output:

```text
10
```

Explanation:

```text
The substrings containing at least one occurrence of a, b, and c are:

"abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc", and "abc" (again).

There are 10 such substrings.
```

### Example 2

Input:

```text
s = "aaacb"
```

Output:

```text
3
```

Explanation:

```text
The valid substrings are "aaacb", "aacb", and "acb".
```

### Example 3

Input:

```text
s = "abc"
```

Output:

```text
1
```

## Constraints

```text
3 <= s.length <= 5 * 10^4
s only consists of a, b or c characters.
```

## Approach

Use a sliding window with counts of `a`, `b`, and `c`.

Expand `right` one character at a time. When the window contains all three characters, every substring that starts anywhere from `left` through `right` and ends at `right` is valid.

There are:

```text
n - right
```

such substrings.

Then shrink from the left while the window remains valid, adding the same count each time.

## Algorithm

1. Initialize `count = [3]int`, `left = 0`, and `result = 0`.
2. Iterate `right` from `0` to `n - 1`.
3. Increase the count of `s[right]`.
4. While the window contains at least one `a`, `b`, and `c`:
   * Add `n - right` to `result`.
   * Decrease the count of `s[left]` and move `left` forward.
5. Return `result`.

## Why This Works

If `[left, right]` is the smallest valid window ending at `right`, then any substring ending at `right` with a start index in:

```text
left, left + 1, ..., right
```

still contains all three characters.

That gives exactly `right - left + 1 = n - right` substrings when counted from the current `left`, but the code adds `n - right` after each successful shrink step, which accumulates all valid windows ending at each `right`.

Equivalently, for each valid window position, all extensions to the right remain valid, so counting by ending index works.

## Complexity Analysis

Let `n` be the length of `s`.

### Time Complexity

```text
O(n)
```

Each pointer moves at most `n` times.

### Space Complexity

```text
O(1)
```

Only a fixed-size count array and a few integers are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `[3]int` is enough because the alphabet size is fixed
* The shrink loop runs only while all three counts are positive
* `result += n - right` counts all valid substrings ending at `right`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* A string missing one required character
* A longer repeated pattern

## Edge Cases

Important cases to consider:

* Minimum length string `"abc"`
* Missing `a`, `b`, or `c` entirely
* Many repeated copies of one character before the others appear
* Multiple valid windows ending at the same position

## Notes

* This is a classic sliding-window counting problem.
* The key insight is that once a window is valid, extending it to the right keeps it valid.
* Because the alphabet has only three characters, an array beats a hash map here.
