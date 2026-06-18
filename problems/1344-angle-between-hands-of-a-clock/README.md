---
id: 1344
title: "Angle Between Hands of a Clock"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/angle-between-hands-of-a-clock/"
contest: "Biweekly Contest 19"
status: "Solved"
language: "Go"
topics:
  - "Math"
go_concepts:
  - "float64 for degree calculations"
  - "math.Abs and math.Min"
  - "Integer modulo for hour normalization"
  - "Floating-point comparison with tolerance in tests"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - math
  - biweekly-contest-19
---

# 1344. Angle Between Hands of a Clock

## Problem Link

LeetCode: `https://leetcode.com/problems/angle-between-hands-of-a-clock/`

## Difficulty

Medium

## Problem Topics

* Math

## What to Know Before Solving

General concepts:

* A clock face has 12 hour marks and 360 degrees in total
* Each hour mark represents 30 degrees
* The minute hand moves 6 degrees per minute
* The hour hand moves continuously, not only on the hour
* The answer must be the smaller angle between the two hands

Go concepts:

* Using `float64` for fractional hand positions and degree results
* Normalizing hour 12 with `hour % 12`
* Using `math.Abs` for absolute difference
* Using `math.Min` to choose the smaller of two angles
* Comparing floating-point results with a small tolerance in tests

## Problem Description

Given two numbers, `hour` and `minutes`, return the smaller angle (in degrees) formed between the hour and the minute hand.

Answers within `10^-5` of the actual value will be accepted as correct.

## Function Signature

Expected LeetCode function signature:

```go
func angleClock(hour int, minutes int) float64 {

}
```

## Examples

### Example 1

Input:

```text
hour = 12, minutes = 30
```

Output:

```text
165
```

### Example 2

Input:

```text
hour = 3, minutes = 30
```

Output:

```text
75
```

### Example 3

Input:

```text
hour = 3, minutes = 15
```

Output:

```text
7.5
```

## Constraints

```text
1 <= hour <= 12
0 <= minutes <= 59
```

## Approach

Convert both hands into positions on the same 12-hour scale, then convert the difference into degrees.

The minute hand position is `minutes / 5`, because 60 minutes map to 12 hour marks.

The hour hand position is `hour % 12 + minutes / 60`, because the hour hand moves with the minutes.

Multiply the absolute difference by 30 to convert hour marks into degrees.

Finally, return the smaller angle between that value and `360 - angle`.

## Algorithm

1. Compute `minuteHandPosition = minutes / 5`.
2. Compute `hourHandPosition = hour % 12 + minutes / 60`.
3. Compute `positionDifference = abs(minuteHandPosition - hourHandPosition)`.
4. Convert to degrees: `angle = positionDifference * 30`.
5. Return `min(angle, 360 - angle)`.

## Why This Works

On a clock:

* Each hour mark is 30 degrees apart
* The minute hand completes 360 degrees in 60 minutes, so it moves 6 degrees per minute
* Dividing minutes by 5 gives the minute hand position in hour-mark units
* The hour hand also moves with minutes, so `minutes / 60` must be added to the hour position

The direct angle and its reflex angle always sum to 360 degrees. Taking the minimum gives the smaller angle required by the problem.

## Complexity Analysis

### Time Complexity

```text
O(1)
```

Only a constant number of arithmetic operations are performed.

### Space Complexity

```text
O(1)
```

No extra data structures are used.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `hour % 12` correctly treats 12 o'clock as position 0
* `minutes / 60` models the hour hand's continuous movement
* `math.Min(angle, 360-angle)` ensures the smaller angle is returned

```go
package anglebetweenhandsofclock

import "math"

func angleClock(hour int, minutes int) float64 {
	minuteHandPosition := float64(minutes) / 5
	hourHandPosition := float64(hour%12) + float64(minutes)/60

	positionDifference := math.Abs(minuteHandPosition - hourHandPosition)
	angle := positionDifference * 30

	return math.Min(angle, 360-angle)
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Hands overlapping at noon
* Hands pointing in opposite directions at 6:00
* A case where only the minute hand moves
* Hour 12 normalized correctly

## Edge Cases

Important cases to consider:

* `hour = 12` and `minutes = 0`
* `hour = 6` and `minutes = 0`
* Small fractional angles such as `3:15`
* Cases where the reflex angle would be smaller without `math.Min`
* `minutes = 0` and `minutes = 59`

## Notes

* This is a pure math problem; no simulation of the clock is needed.
* Always remember that the hour hand moves with minutes.
* Because the answer is a `float64`, tests should compare with a small tolerance such as `1e-5`.
