---
id: 1732
title: "Find the Highest Altitude"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/find-the-highest-altitude/"
contest: "Biweekly Contest 44"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Prefix Sum"
go_concepts:
  - "Slice iteration with range"
  - "Running sum tracking"
  - "Tracking a maximum while scanning"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - prefix-sum
  - biweekly-contest-44
---

# 1732. Find the Highest Altitude

## Problem Link

LeetCode: `https://leetcode.com/problems/find-the-highest-altitude/`

## Difficulty

Easy

## Problem Topics

* Array
* Prefix Sum

## What to Know Before Solving

General concepts:

* Prefix sums represent cumulative totals after each step
* Starting altitude is 0 at point 0
* Each `gain[i]` changes altitude between consecutive points
* The answer is the maximum altitude reached at any point

Go concepts:

* Iterating over a slice with `for _, value := range gain`
* Maintaining running totals with simple integer variables
* Updating a maximum while scanning
* Writing table-driven tests with the `testing` package

## Problem Description

There is a biker going on a road trip. The road trip consists of `n + 1` points at different altitudes. The biker starts his trip on point 0 with altitude equal to 0.

You are given an integer array `gain` of length `n` where `gain[i]` is the net gain in altitude between points `i` and `i + 1` for all `(0 <= i < n)`.

Return the highest altitude of a point.

## Function Signature

Expected LeetCode function signature:

```go
func largestAltitude(gain []int) int {

}
```

## Examples

### Example 1

Input:

```text
gain = [-5,1,5,0,-7]
```

Output:

```text
1
```

Explanation:

```text
The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
```

### Example 2

Input:

```text
gain = [-4,-3,-2,-1,4,3,2]
```

Output:

```text
0
```

Explanation:

```text
The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
```

## Constraints

```text
n == gain.length
1 <= n <= 100
-100 <= gain[i] <= 100
```

## Approach

Track the current altitude while moving through the trip.

Start at altitude 0. For each gain value, add it to the current altitude and update the highest altitude seen so far.

This is a prefix-sum style scan where we only need the running maximum, not the full altitude array.

## Algorithm

1. Set `currentAltitude = 0` and `highestAltitude = 0`.
2. Iterate over each value in `gain`.
3. Add the value to `currentAltitude`.
4. If `currentAltitude > highestAltitude`, update `highestAltitude`.
5. Return `highestAltitude`.

## Why This Works

The biker starts at altitude 0. After processing the first `i` gain values, the altitude at point `i + 1` is exactly the prefix sum of those gains.

By updating the maximum after every step, we visit every reachable altitude without storing the full sequence. Since the highest point must appear at one of these prefix-sum states, the maximum we track is the correct answer.

## Complexity Analysis

Let `n` be the length of `gain`.

### Time Complexity

```text
O(n)
```

Each gain value is processed once.

### Space Complexity

```text
O(1)
```

Only two integer variables are used besides the input.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `highestAltitude` is initialized to 0 because the starting point counts
* The loop uses `range` over the slice values directly
* No extra slice is needed to store all altitudes

```go
package findthehighestaltitude

func largestAltitude(gain []int) int {
	currentAltitude := 0
	highestAltitude := 0

	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain

		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Single-element gain arrays
* All-positive gains
* A peak in the middle of the trip
* Cases where the highest altitude remains the starting value 0

## Edge Cases

Important cases to consider:

* A single gain value
* All negative gains
* All positive gains
* Highest altitude reached at the first point after start
* Highest altitude reached at the last point
* Zero gains mixed with positive and negative values

## Notes

* This is a classic prefix-sum maximum problem.
* Initializing `highestAltitude` to 0 is important because the answer may never exceed the starting altitude.
* A slightly more compact version could use `max(highestAltitude, currentAltitude)` if using Go 1.21+ built-in `max`.
