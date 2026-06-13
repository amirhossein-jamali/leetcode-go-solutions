# 3838. Weighted Word Mapping

## Problem Link

LeetCode: `https://leetcode.com/problems/weighted-word-mapping/`

## Difficulty

Easy

## Problem Description

You are given an array of strings `words`, where each string contains lowercase English letters.

You are also given an integer array `weights` of length `26`, where `weights[i]` represents the weight of the `i`th lowercase English letter.

The weight of a word is defined as the sum of the weights of all characters in that word.

For each word:

1. Compute the total weight of the word.
2. Take the total weight modulo `26`.
3. Map the modulo result to a lowercase English letter using reverse alphabetical order:

```text
0  -> 'z'
1  -> 'y'
...
25 -> 'a'
```

Return a string formed by concatenating the mapped characters for all words in order.

## Examples

### Example 1

Input:

```text
words = ["abcd","def","xyz"]
weights = [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
```

Output:

```text
"rij"
```

Explanation:

```text
"abcd" has weight 5 + 3 + 12 + 14 = 34.
34 % 26 = 8, which maps to 'r'.

"def" has weight 14 + 1 + 2 = 17.
17 % 26 = 17, which maps to 'i'.

"xyz" has weight 7 + 7 + 2 = 16.
16 % 26 = 16, which maps to 'j'.
```

Therefore, the final result is:

```text
"rij"
```

### Example 2

Input:

```text
words = ["a","b","c"]
weights = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
```

Output:

```text
"yyy"
```

Explanation:

```text
Each word has weight 1.
1 % 26 = 1, which maps to 'y'.
```

Therefore, the final result is:

```text
"yyy"
```

### Example 3

Input:

```text
words = ["abcd"]
weights = [7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5]
```

Output:

```text
"g"
```

Explanation:

```text
"abcd" has weight 7 + 5 + 3 + 4 = 19.
19 % 26 = 19, which maps to 'g'.
```

Therefore, the final result is:

```text
"g"
```

## Constraints

```text
1 <= words.length <= 100
1 <= words[i].length <= 10
weights.length == 26
1 <= weights[i] <= 100
words[i] consists of lowercase English letters.
```

## Approach

Each lowercase English letter can be mapped to an index from `0` to `25`.

Since `weights[0]` belongs to `'a'`, `weights[1]` belongs to `'b'`, and so on, the index of a character can be calculated as:

```text
character - 'a'
```

For each word, we sum the weights of its characters.

After computing the total weight, we take it modulo `26`. This gives a value between `0` and `25`.

The problem maps this value using reverse alphabetical order:

```text
0 -> 'z'
1 -> 'y'
...
25 -> 'a'
```

So the mapped character can be calculated directly as:

```text
'z' - modulo_result
```

Because the output contains exactly one character per word, we can allocate a byte slice with length equal to `len(words)` and fill each position directly.

## Algorithm

1. Create a byte slice `result` with length equal to the number of words.
2. Iterate over `words` with index `i`.
3. For each word:

    1. Initialize `sum` to `0`.
    2. Iterate over every character in the word.
    3. Convert the character to its alphabet index using `word[j] - 'a'`.
    4. Add the corresponding value from `weights` to `sum`.
    5. Compute `sum % 26`.
    6. Convert the modulo result to a reverse alphabetical character using `'z' - modulo`.
    7. Store the mapped character at `result[i]`.
4. Convert `result` to a string and return it.

## Why This Works

Each word's weight is defined as the sum of the weights of its characters, so iterating through the characters and accumulating their corresponding weights computes the correct total.

Subtracting `'a'` from a lowercase English letter gives its zero-based alphabet index. This index matches the position of that letter inside the `weights` array.

Taking the sum modulo `26` produces a value in the range `0` to `25`, which exactly matches the number of lowercase English letters.

Since the mapping is in reverse alphabetical order, subtracting the modulo result from `'z'` gives the required output character.

Therefore, each word is converted to the correct mapped character, and storing these characters in order produces the final answer.

## Complexity Analysis

Let `n` be the number of words and let `L` be the total number of characters across all words.

### Time Complexity

```text
O(L)
```

Every character is visited exactly once.

If `m` is the maximum length of a word, then `L <= n * m`, so the upper bound can also be written as:

```text
O(n * m)
```

### Space Complexity

```text
O(n)
```

The output contains one character for each word.

Excluding the output, the auxiliary space usage is:

```text
O(1)
```

## Code

The Go solution is available in:

```text
solution.go
```

A memory-efficient implementation uses a preallocated byte slice:

```go
package weightedwordmapping

func mapWordWeights(words []string, weights []int) string {
	result := make([]byte, len(words))

	for i, word := range words {
		sum := 0

		for j := 0; j < len(word); j++ {
			sum += weights[int(word[j]-'a')]
		}

		result[i] = byte('z' - (sum % 26))
	}

	return string(result)
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

## Edge Cases

Important cases to consider:

* A single word
* A single-letter word
* Repeated letters inside a word
* A word whose total weight is divisible by `26`
* All weights being the same
* The modulo result being `0`, which should map to `'z'`
* The modulo result being `25`, which should map to `'a'`
