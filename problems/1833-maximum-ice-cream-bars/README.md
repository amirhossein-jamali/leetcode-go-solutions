---
id: 1833
title: "Maximum Ice Cream Bars"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/maximum-ice-cream-bars/"
contest: "Weekly Contest 237"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Greedy"
  - "Sorting"
  - "Counting Sort"
go_concepts:
  - "Slice iteration with range"
  - "Frequency counting with a slice"
  - "Greedy purchase from smallest price upward"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - greedy
  - counting-sort
  - weekly-contest-237
---

# 1833. Maximum Ice Cream Bars

## Problem Link

LeetCode: `https://leetcode.com/problems/maximum-ice-cream-bars/`

## Difficulty

Medium

## Problem Topics

* Array
* Greedy
* Sorting
* Counting Sort

## What to Know Before Solving

General concepts:

* To maximize the number of items bought with fixed coins, buy the cheapest items first
* Counting sort works well when values are bounded and repeated
* A frequency array maps each price to how many bars cost that price
* Greedy choice: always take as many of the current cheapest available price as possible

Go concepts:

* Iterating over a slice with `for _, value := range costs`
* Building a frequency slice with `make([]int, maxCost+1)`
* Integer division to compute how many items fit in remaining coins
* Writing table-driven tests with the `testing` package

## Problem Description

It is a sweltering summer day, and a boy wants to buy some ice cream bars.

At the store, there are `n` ice cream bars. You are given an integer array `costs` of length `n` where `costs[i]` is the price of the `i`th ice cream bar in coins. The boy initially has `coins` coins to spend, and he wants to buy as many ice cream bars as possible.

Note: The boy can buy the ice cream bars in any order.

Return the maximum number of ice cream bars the boy can buy with `coins` coins.

You must solve the problem by counting sort.

## Function Signature

Expected LeetCode function signature:

```go
func maxIceCream(costs []int, coins int) int {

}
```

## Examples

### Example 1

Input:

```text
costs = [1,3,2,4,1]
coins = 7
```

Output:

```text
4
```

Explanation:

```text
The boy can buy ice cream bars at indices 0, 1, 2, and 4 for a total price of 1 + 3 + 2 + 1 = 7.
```

### Example 2

Input:

```text
costs = [10,6,8,7,7,8]
coins = 5
```

Output:

```text
0
```

Explanation:

```text
The boy cannot afford any of the ice cream bars.
```

### Example 3

Input:

```text
costs = [1,6,3,1,2,5]
coins = 20
```

Output:

```text
6
```

Explanation:

```text
The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.
```

## Constraints

```text
costs.length == n
1 <= n <= 10^5
1 <= costs[i] <= 10^5
1 <= coins <= 10^8
```

## Approach

Use counting sort to group ice cream bars by price, then buy greedily from the cheapest price upward.

First, count how many bars exist at each price. Then scan prices from `1` to the maximum cost. At each price, buy as many bars as possible with the remaining coins before moving to the next price.

This avoids sorting the full input array explicitly while still processing prices in ascending order.

## Algorithm

1. Find `maxCost`, the largest value in `costs`.
2. Build `frequency`, where `frequency[price]` is the number of bars with that price.
3. Set `bought = 0`.
4. For each `price` from `1` to `maxCost` while `coins >= price`:
   - Skip prices with zero frequency.
   - Compute `count = coins / price`.
   - Cap `count` at `frequency[price]`.
   - Add `count` to `bought` and subtract `count * price` from `coins`.
   - If not all bars at this price were bought, stop because remaining coins cannot afford a higher price.
5. Return `bought`.

## Why This Works

Buying cheaper bars first never reduces the total number of bars you can afford. If you can buy a bar of price `p`, using those coins on a bar of price `p` is at least as good as saving them for a more expensive bar, because each bar counts as one item regardless of price.

Counting sort lets us visit prices in ascending order without sorting all `n` values with a comparison sort. By taking as many bars as possible at each price before moving upward, we maximize the count for the remaining coins.

## Complexity Analysis

Let `n` be the length of `costs` and `m` be the maximum value in `costs`.

### Time Complexity

```text
O(n + m)
```

Finding `maxCost` and building the frequency array each scan `costs` once. Then we iterate through prices from `1` to `m`.

### Space Complexity

```text
O(m)
```

The frequency slice stores one count per possible price up to `maxCost`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `frequency` is sized as `maxCost+1` so each price index is valid
* The loop stops early when remaining coins cannot buy all bars at the current price
* Integer division gives the maximum number of bars affordable at the current price

```go
package maximumicecreambars

func maxIceCream(costs []int, coins int) int {
	maxCost := 0

	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	frequency := make([]int, maxCost+1)

	for _, cost := range costs {
		frequency[cost]++
	}

	bought := 0

	for price := 1; price <= maxCost && coins >= price; price++ {
		if frequency[price] == 0 {
			continue
		}

		count := coins / price

		if count > frequency[price] {
			count = frequency[price]
		}

		bought += count
		coins -= count * price

		if count < frequency[price] {
			break
		}
	}

	return bought
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* A single affordable bar
* A single unaffordable bar
* Multiple bars with the same price
* Buying only the cheapest repeated prices

## Edge Cases

Important cases to consider:

* No bar is affordable
* All bars are affordable
* Many bars share the same price
* Remaining coins are not enough for the next price
* Coins exactly match the total cost of all bought bars

## Notes

* This is a classic greedy plus counting sort problem.
* An alternative is to sort `costs` and scan from smallest to largest, but counting sort fits the bounded price range well.
* The early `break` is important: once you cannot buy every bar at the current price, you also cannot afford any higher price.
