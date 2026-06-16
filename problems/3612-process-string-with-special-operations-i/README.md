---
id: 3612
title: "Process String with Special Operations I"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/process-string-with-special-operations-i/"
contest: "Weekly Contest 458"
status: "Solved"
language: "Go"
topics:
  - "String"
  - "Simulation"
go_concepts:
  - "Byte slices as mutable string builders"
  - "Switch on byte characters"
  - "Slice truncation with reslicing"
  - "append with slice spread"
  - "In-place reversal with two pointers"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - string
  - simulation
  - weekly-contest-458
---

# 3612. Process String with Special Operations I

## Problem Link

LeetCode: `https://leetcode.com/problems/process-string-with-special-operations-i/`

## Difficulty

Medium

## Problem Topics

* String
* Simulation

## What to Know Before Solving

General concepts:

* How to simulate a sequence of operations from left to right
* How string-building operations like append, delete, duplicate, and reverse interact
* Why order matters when operations can change the current result length

Go concepts:

* Using `[]byte` as a mutable buffer for building a result string
* Iterating over a `string` by byte index for ASCII input
* Truncating a slice with `result[:len(result)-1]`
* Duplicating a slice with `append(result, result...)`
* Reversing a slice in place with two pointers
* Handling special characters with a `switch` statement
* Writing table-driven tests with the `testing` package

## Problem Description

You are given a string `s` consisting of lowercase English letters and the special characters: `*`, `#`, and `%`.

Build a new string `result` by processing `s` according to the following rules from left to right:

* If the character is a lowercase English letter, append it to `result`.
* A `*` removes the last character from `result`, if it exists.
* A `#` duplicates the current `result` and appends it to itself.
* A `%` reverses the current `result`.

Return the final string `result` after processing all characters in `s`.

## Function Signature

Expected LeetCode function signature:

```go
func processStr(s string) string {

}
```

## Examples

### Example 1

Input:

```text
s = "a#b%*"
```

Output:

```text
"ba"
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

Thus, the final result is `"ba"`.

### Example 2

Input:

```text
s = "z*#"
```

Output:

```text
""
```

Explanation:

```text
i  s[i]  Operation        Current result
0  'z'   Append 'z'       "z"
1  '*'   Remove last char ""
2  '#'   Duplicate result ""
```

Thus, the final result is `""`.

## Constraints

```text
1 <= s.length <= 20
s consists of only lowercase English letters and special characters *, #, and %.
```

## Approach

We maintain a mutable byte slice called `result` and process the input string from left to right.

Each character triggers one of four actions:

* A letter is appended directly.
* `*` removes the last byte if the slice is not empty.
* `#` appends a copy of the current slice to itself.
* `%` reverses the current slice in place.

Because every operation only depends on the current `result`, a direct left-to-right simulation is enough.

## Algorithm

1. Initialize `result` as an empty `[]byte`.
2. Iterate over every character in `s`.
3. For each character:
   * If it is `*`, remove the last element when `len(result) > 0`.
   * If it is `#`, append the current slice to itself.
   * If it is `%`, swap characters from both ends until the slice is reversed.
   * Otherwise, append the letter to `result`.
4. Convert `result` to a string and return it.

## Why This Works

The problem defines the result after each character using only the current `result` state.

Processing characters in order means each operation is applied exactly when the problem expects it. There is no need to look ahead or backtrack, because later operations always use the updated `result`.

Each rule is implemented directly:

* Append adds one character.
* Delete shortens the slice by one when possible.
* Duplicate doubles the current content.
* Reverse changes the order of the existing content.

Therefore, after the last character is processed, `result` matches the required final string.

## Complexity Analysis

Let `n` be the length of `s`, and let `m` be the maximum length of `result` during processing.

### Time Complexity

```text
O(n * m)
```

In the worst case, operations like `#` and `%` can touch most or all of the current result. With the given constraint `n <= 20`, this simulation is easily fast enough.

### Space Complexity

```text
O(m)
```

The byte slice stores the current result. Excluding the output, auxiliary space is also `O(m)`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `[]byte` is used because the result must be modified in place.
* `append(result, result...)` cleanly duplicates the current slice for the `#` operation.
* Reversal is done in place with two indices, so no extra buffer is needed.

```go
package processstringwithspecialoperationsi

func processStr(s string) string {
	result := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}

		case '#':
			result = append(result, result...)

		case '%':
			for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
				result[left], result[right] = result[right], result[left]
			}

		default:
			result = append(result, ch)
		}
	}

	return string(result)
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Single-letter input
* Duplicate-only and reverse-only operations
* No-op delete, duplicate, and reverse on an empty result
* Combined operation sequences

## Edge Cases

Important cases to consider:

* `*` on an empty result should do nothing
* `#` on an empty result should stay empty
* `%` on an empty result should stay empty
* A single character input with no special operations
* Operations that shrink the result before a later duplicate or reverse
* Chains where duplicate and reverse interact, such as `"a#%"`

## Notes

* This is a straightforward simulation problem; the main challenge is applying each rule correctly in order.
* Using `[]byte` instead of repeated string concatenation avoids unnecessary allocations.
* For `#`, `append(result, result...)` is idiomatic Go for appending a slice copy to itself.
