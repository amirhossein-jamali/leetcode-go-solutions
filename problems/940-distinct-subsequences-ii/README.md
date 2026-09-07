---
id: 940
title: "Distinct Subsequences II"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/distinct-subsequences-ii/"
contest: "Weekly Contest 110"
status: "Solved"
language: "Go"
topics:
  - "String"
  - "Dynamic Programming"
go_concepts:
  - "Fixed-size array for last-ending counts"
  - "Modulo arithmetic with int64"
  - "Incremental subsequence counting"
  - "Byte indexing with s[i]-'a'"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - string
  - dynamic-programming
  - weekly-contest-110
---

# 940. Distinct Subsequences II

## Problem Link

LeetCode: `https://leetcode.com/problems/distinct-subsequences-ii/`

## Difficulty

Hard

## Problem Topics

* String
* Dynamic Programming

## What to Know Before Solving

General concepts:

* A subsequence keeps relative order but may skip characters
* Repeated letters create duplicate subsequences unless they are removed carefully
* If `total` is the number of distinct non-empty subsequences so far, appending a new unique character roughly doubles the set and adds one more subsequence consisting of that character alone
* When the same character appears again, subtract the old count of subsequences that ended with that character to avoid duplicates

Go concepts:

* Using `[26]int64` to track how many distinct subsequences end with each letter
* Updating a running total with modulo `1_000_000_007`
* Normalizing negative or oversized modulo results manually
* Mapping characters with `s[i] - 'a'`
* Writing table-driven tests with the `testing` package

## Problem Description

Given a string `s`, return the number of distinct non-empty subsequences of `s`.

Since the answer may be very large, return it modulo `10^9 + 7`.

## Function Signature

Expected LeetCode function signature:

```go
func distinctSubseqII(s string) int {

}
```

## Examples

### Example 1

Input:

```text
s = "abc"
```

Output:

```text
7
```

Explanation:

```text
The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".
```

### Example 2

Input:

```text
s = "aba"
```

Output:

```text
6
```

Explanation:

```text
The 6 distinct subsequences are "a", "b", "ab", "aa", "ba", and "aba".
```

### Example 3

Input:

```text
s = "aaa"
```

Output:

```text
3
```

Explanation:

```text
The 3 distinct subsequences are "a", "aa", and "aaa".
```

## Constraints

```text
1 <= s.length <= 2000
s consists of lowercase English letters
```

## Approach

Maintain:

* `total`: number of distinct non-empty subsequences seen so far
* `end[c]`: number of distinct subsequences that end with character `c`

When processing `s[i]` with character `c`:

1. Every existing subsequence can optionally include this new character, which doubles the count.
2. The single-character subsequence `"c"` itself is also new, so add `1`.
3. Some new subsequences ending in `c` were already counted the last time we saw `c`, so subtract `end[c]`.

That gives:

```text
total = total * 2 + 1 - end[c]
```

Then update `end[c]` to the new number of distinct subsequences ending in `c`.

## Algorithm

1. Initialize `total = 0` and `end` as a zero array of size `26`.
2. For each character `s[i]`:
   * Let `c = s[i] - 'a'` and save `old = end[c]`.
   * Update `total = (total * 2 + 1 - old) mod (10^9 + 7)`.
   * Set `end[c] = (total + 1) mod (10^9 + 7)` in this implementation's bookkeeping form.
3. Return `total`.

## Why This Works

### Why doubling works

Suppose we already have a set of distinct subsequences. When a new character arrives, every old subsequence can either:

* stay unchanged, or
* extend by appending the new character

That creates a doubled set, plus one extra subsequence made only of the new character.

### Why we subtract `end[c]`

If character `c` appeared before, some subsequences ending in `c` were already counted in previous steps. Appending the current `c` again would recreate duplicates such as choosing an earlier `c` versus this one.

Subtracting `end[c]` removes exactly those duplicate endings.

### Why modulo is needed

The number of subsequences grows exponentially, so every update must stay inside modulo `10^9 + 7`.

## Complexity Analysis

Let:

```text
n = len(s)
```

### Time Complexity

```text
O(n)
```

We scan the string once and do constant work per character.

### Space Complexity

```text
O(1)
```

Only `total` and a fixed array of size `26` are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Use `int64` for intermediate modulo arithmetic
* Keep `end[c]` as the count of distinct subsequences ending with character `c`
* Normalize `total` after subtraction because `total * 2 + 1 - old` can become negative before modulo

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Single-character input
* Two distinct characters
* Repeated letters with partial overlap

## Edge Cases

Important cases to consider:

* `s = "a"` with answer `1`
* Repeated characters such as `"aaa"`
* Alternating repeats such as `"aba"`
* Long strings where modulo arithmetic matters
* Strings where many characters are identical

## Notes

* This problem is related to LeetCode 115 and 940-style counting, but here the target is not fixed.
* The `end[c]` array is the key to deduplicating subsequences that end with the same character.
* A 2D DP solution also works, but the optimized counting form is shorter and faster.
