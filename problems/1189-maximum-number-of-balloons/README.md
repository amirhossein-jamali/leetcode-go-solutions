---
id: 1189
title: "Maximum Number of Balloons"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-number-of-balloons/"
contest: "Weekly Contest 154"
status: "Solved"
language: "Go"
topics:
  - "Hash Table"
  - "String"
  - "Counting"
go_concepts:
  - "String iteration by byte index"
  - "Switch on byte values"
  - "Character frequency counting"
  - "Taking the minimum across counts"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - hash-table
  - string
  - counting
  - weekly-contest-154
---

# 1189. Maximum Number of Balloons

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-number-of-balloons/`

## Difficulty

Easy

## Problem Topics

* Hash Table
* String
* Counting

## What to Know Before Solving

General concepts:

* Count how many times each required letter appears
* The word `balloon` needs two `l` characters and two `o` characters
* The answer is limited by the scarcest required letter
* Extra letters that are not part of `balloon` can be ignored

Go concepts:

* Iterating over a string with a byte index
* Using `switch` on `text[i]`
* Updating integer counters
* Comparing values to find a minimum

## Problem Description

Given a string `text`, use the characters of `text` to form as many instances of the word `balloon` as possible.

Each character in `text` can be used at most once.

Return the maximum number of instances that can be formed.

## Function Signature

Expected LeetCode function signature:

```go
func maxNumberOfBalloons(text string) int {

}
```

## Examples

### Example 1

Input:

```text
text = "nlaebolko"
```

Output:

```text
1
```

Explanation:

```text
One instance of "balloon" can be formed.
```

### Example 2

Input:

```text
text = "loonbalxballpoon"
```

Output:

```text
2
```

Explanation:

```text
Two instances of "balloon" can be formed.
```

### Example 3

Input:

```text
text = "leetcode"
```

Output:

```text
0
```

Explanation:

```text
No instance of "balloon" can be formed.
```

## Constraints

```text
1 <= text.length <= 10^4
text consists of lowercase English letters only
```

## Approach

Count how many times each letter needed for `balloon` appears in `text`.

The required letters are:

```text
b: 1
a: 1
l: 2
o: 2
n: 1
```

After counting, divide the `l` and `o` counts by 2 because each `balloon` needs two of each.

The answer is the minimum of the five resulting counts.

## Algorithm

1. Initialize counters for `b`, `a`, `l`, `o`, and `n` to 0.
2. Scan `text` once and increment the matching counter for each relevant letter.
3. Divide the `l` and `o` counters by 2.
4. Start with `result = b`.
5. Update `result` to the minimum of `result`, `a`, `l`, `o`, and `n`.
6. Return `result`.

## Why This Works

Each `balloon` consumes exactly one `b`, one `a`, two `l`, two `o`, and one `n`.

If `text` contains `count(b)` copies of `b`, then at most `count(b)` balloons can be formed because every balloon needs one `b`. The same logic applies to the other letters.

For `l` and `o`, each balloon needs two copies, so the number of balloons supported by those letters is `count(l) / 2` and `count(o) / 2`.

The final answer must satisfy all five letter requirements at the same time, so it is the minimum of those five values.

## Complexity Analysis

Let `n` be the length of `text`.

### Time Complexity

```text
O(n)
```

The string is scanned once.

### Space Complexity

```text
O(1)
```

Only a fixed number of integer counters are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Only letters in `balloon` are counted
* `l` and `o` are halved before the minimum step
* The minimum is computed with simple comparisons instead of extra data structures

```go
package maximumnumberofballoons

func maxNumberOfBalloons(text string) int {
	var b, a, l, o, n int

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	l /= 2
	o /= 2

	result := b

	if a < result {
		result = a
	}
	if l < result {
		result = l
	}
	if o < result {
		result = o
	}
	if n < result {
		result = n
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
* One exact `balloon`
* Two complete `balloon` strings
* Missing required letters
* Not enough `l` characters for a full word
* A single-character input
* Extra unrelated letters

## Edge Cases

Important cases to consider:

* No usable letters for `balloon`
* Exactly one complete `balloon`
* Multiple complete `balloon` strings concatenated
* Missing one required letter
* Only one `l` or one `o`
* Many extra letters that are not part of `balloon`

## Notes

* This problem is equivalent to LeetCode 2287: Rearrange Characters to Make Target String.
* A hash map also works, but five integer counters are enough because the target word is fixed.
* Letters outside `balloon` never affect the answer.
