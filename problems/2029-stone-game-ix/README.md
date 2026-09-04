---
id: 2029
title: "Stone Game IX"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/stone-game-ix/"
contest: "Weekly Contest 261"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Math"
  - "Greedy"
  - "Counting"
  - "Game Theory"
go_concepts:
  - "Fixed-size array [3]int"
  - "Integer modulo"
  - "Absolute difference without math.Abs"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - math
  - game-theory
  - counting
  - weekly-contest-261
---

# 2029. Stone Game IX

## Problem Link

LeetCode: `https://leetcode.com/problems/stone-game-ix/`

## Difficulty

Medium

## Problem Topics

* Array
* Math
* Greedy
* Counting
* Game Theory

## What to Know Before Solving

General concepts:

* A player loses by making the removed-stone sum divisible by `3`
* If the stones run out without that happening, Bob wins
* Only each stone's value modulo `3` affects the sum
* Stones of type `0` do not change the current remainder

Go concepts:

* Counting remainders in a fixed `[3]int` array
* Classifying stones with `stone % 3`
* Computing an absolute difference without importing `math`
* Writing table-driven tests with the `testing` package

## Problem Description

Alice and Bob take turns removing one stone, and Alice starts.

A player loses immediately if the sum of all removed stones is divisible by `3`. If no stones remain and nobody has lost that way, Bob wins.

Return `true` if Alice wins when both play optimally.

## Function Signature

Expected LeetCode function signature:

```go
func stoneGameIX(stones []int) bool {

}
```

## Examples

### Example 1

Input:

```text
stones = [2,1]
```

Output:

```text
true
```

### Example 2

Input:

```text
stones = [2]
```

Output:

```text
false
```

### Example 3

Input:

```text
stones = [5,1,2,4,3]
```

Output:

```text
false
```

## Constraints

```text
1 <= stones.length <= 10^5
1 <= stones[i] <= 10^4
```

## Approach

Only the values modulo `3` matter, so count stones into `cnt[0]`, `cnt[1]`, and `cnt[2]`.

Type `0` stones do not change the current remainder. They only flip who is forced to break the `1/2` chain. That reduces the game to two cases:

If `cnt[0]` is even:

```text
Alice wins iff cnt[1] > 0 && cnt[2] > 0
```

If `cnt[0]` is odd:

```text
Alice wins iff abs(cnt[1] - cnt[2]) > 2
```

## Algorithm

1. Count each stone into `cnt[stone % 3]`.
2. If `cnt[0]` is even, return whether both `cnt[1]` and `cnt[2]` are positive.
3. Otherwise return whether `abs(cnt[1] - cnt[2]) > 2`.

## Why This Works

Alice cannot open with a type `0` stone. After she opens with `1` or `2`, the non-zero stones must alternate `1, 2, 1, 2, ...` (or the symmetric sequence). A type `0` stone can be inserted without changing the remainder, which only swaps the player who would next be forced off that chain.

When `cnt[0]` is even, Bob can always answer a type `0` stone with another type `0` stone. Alice then wins only if both non-zero types exist, so she can start a chain that eventually forces Bob to make the sum divisible by `3`.

When `cnt[0]` is odd, Alice can use the extra type `0` stone to flip that parity. She still needs a long enough imbalance between `cnt[1]` and `cnt[2]`; the threshold is `abs(cnt[1] - cnt[2]) > 2`.

If the stones are exhausted without anyone hitting a multiple of `3`, Bob wins. That is already covered by the two formulas.

## Complexity Analysis

Let:

```text
n = len(stones)
```

### Time Complexity

```text
O(n)
```

The array is scanned once. The remaining arithmetic is constant time.

### Space Complexity

```text
O(1)
```

Only a fixed `[3]int` counter is used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Remainders are stored in a stack-allocated `[3]int`
* The odd-`cnt[0]` branch computes the absolute difference without `math.Abs`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* A single multiple of `3`
* Stones of only one non-zero remainder
* Even `cnt[0]` with both remainders present
* Odd `cnt[0]` on both sides of the `abs(cnt[1] - cnt[2]) > 2` threshold

## Edge Cases

Important cases to consider:

* `n = 1`
* All stones divisible by `3`
* Only remainder `1` or only remainder `2`
* `cnt[0]` even vs odd
* `abs(cnt[1] - cnt[2])` equal to `2` vs greater than `2`

## Notes

* Simulating the game turn by turn is too slow for `n <= 10^5`.
* After the remainder counts are known, the winner is decided in `O(1)`.
