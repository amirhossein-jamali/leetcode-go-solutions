---
id: 1386
title: "Cinema Seat Allocation"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/cinema-seat-allocation/"
contest: "Biweekly Contest 22"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
  - "Greedy"
  - "Bit Manipulation"
go_concepts:
  - "map[int]uint8 for sparse row occupancy"
  - "Bitwise OR to mark reserved seats"
  - "Bitwise AND against family-block masks"
  - "Unsigned comparison to drop seats 1 and 10"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - hash-table
  - greedy
  - bit-manipulation
  - biweekly-contest-22
---

# 1386. Cinema Seat Allocation

## Problem Link

LeetCode: `https://leetcode.com/problems/cinema-seat-allocation/`

## Difficulty

Medium

## Problem Topics

* Array
* Hash Table
* Greedy
* Bit Manipulation

## What to Know Before Solving

General concepts:

* A cinema row has 10 seats, but a four-person family can occupy only one of three blocks: `2..5`, `4..7`, or `6..9`
* Seats `1` and `10` sit outside every family block
* Each row can hold at most two families, in the non-overlapping left and right blocks
* `n` can be as large as `10^9`, so the algorithm must not scan every row

Go concepts:

* Storing sparse row occupancy in `map[int]uint8`
* Setting bits with `|=` and testing blocks with `&`
* Using `uint(seat) < 8` to reject seats outside `2..9`
* Writing table-driven tests with the `testing` package

## Problem Description

A cinema has `n` rows of seats, numbered from `1` to `n`. Each row has 10 seats, numbered from `1` to `10`.

You are given a 2D integer array `reservedSeats`, where `reservedSeats[i] = [rowi, seati]` means that seat `seati` in row `rowi` is already reserved.

A four-person family occupies four adjacent seats in the same row. The only valid blocks are:

* seats `2, 3, 4, 5`
* seats `4, 5, 6, 7`
* seats `6, 7, 8, 9`

A block can be used only if none of its seats are reserved. Each seat can belong to at most one family.

Return the maximum number of four-person families that can be allocated.

## Function Signature

Expected LeetCode function signature:

```go
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
```

Output:

```text
4
```

Explanation:

```text
Row 1 can seat one family in the middle block.
Row 2 can seat one family in the left block.
Row 3 has only aisle seats reserved, so it can seat two families.
```

### Example 2

Input:

```text
n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
```

Output:

```text
2
```

### Example 3

Input:

```text
n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
```

Output:

```text
4
```

## Constraints

```text
1 <= n <= 10^9
1 <= reservedSeats.length <= min(10 * n, 10^4)
reservedSeats[i].length == 2
1 <= reservedSeats[i][0] <= n
1 <= reservedSeats[i][1] <= 10
All reservedSeats[i] are distinct.
```

## Approach

Only seats `2` through `9` affect family placement. Seats `1` and `10` never belong to a valid block, so they can be ignored.

Store occupancy for rows that have at least one reserved seat in `2..9` in a hash map `map[int]uint8`. After shifting each seat by `seat - 2`, seats `2..9` occupy bits `0..7`:

```text
seat:  2  3  4  5  6  7  8  9
bit:   0  1  2  3  4  5  6  7
```

The three family blocks become constant bitmasks:

* `0x0F` (`00001111`) covers seats `2..5`
* `0x3C` (`00111100`) covers seats `4..7`
* `0xF0` (`11110000`) covers seats `6..9`

Start from the empty-cinema capacity `2 * n`. Every mapped row has lost at least one of the two independent side blocks, so subtract `len(rows)`. If a mapped row also intersects all three masks, it has no remaining valid block, so subtract one more.

## Algorithm

1. Create `rows := make(map[int]uint8, len(reservedSeats))`.
2. For each `[row, seat]` in `reservedSeats`:
   * Compute `seat - 2`.
   * If `uint(seat - 2) < 8`, set bit `seat - 2` in `rows[row]`.
3. Set `result = 2 * n - len(rows)`.
4. For each occupancy mask in `rows`:
   * If it intersects `0x0F`, `0x3C`, and `0xF0`, decrement `result`.
5. Return `result`.

## Why This Works

### Empty rows hold two families

A row with no reserved seat in `2..9` has both non-overlapping blocks `2..5` and `6..9` free. Those two families do not share a seat, so the row contributes `2`. Using the overlapping middle block `4..7` would only seat one family, which is worse.

### A reservation in `2..9` drops the row to at most one family

Every seat in `2..9` belongs to at least one of the two side blocks:

* seats `2..5` block the left family
* seats `6..9` block the right family

Therefore any mapped row has already lost at least one of the two independent families. Its capacity is at most `1`. Subtracting `len(rows)` from `2 * n` encodes this: untouched rows still contribute `2`, and each mapped row is assumed to contribute `1`.

### All three masks occupied means zero families

If the occupancy intersects `0x0F`, `0x3C`, and `0xF0`, then the left, middle, and right blocks are all blocked. The row contributes `0`, so one more unit is subtracted from the `2 * n - len(rows)` baseline.

### At least one free mask means exactly one family

If any of the three masks has no reserved bit, that block is free. Combined with the previous point that a mapped row cannot seat two families, the capacity is exactly `1`. No extra subtraction is needed.

### Rows are independent

A family must sit in a single row, and seats are never shared across rows. The maximum for the cinema is the sum of the per-row maxima.

## Complexity Analysis

Let:

```text
m = len(reservedSeats)
r = number of rows with at least one reserved seat in 2..9
```

Then `r <= m`, and `m <= 10^4`.

### Time Complexity

```text
O(m)
```

Each reserved seat is processed once. The second loop visits only the `r` mapped rows.

Scanning all `n` rows is impossible: `n` can be `10^9`.

`O(m)` is asymptotically optimal. Every reserved seat in `2..9` can change the answer, so any correct algorithm must read the whole `reservedSeats` array. Linear in the input size is the best possible.

### Space Complexity

```text
O(r)
```

The hash map stores one `uint8` per row that has an effective reservation. Seats `1` and `10` never create a map entry.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `uint(seat) < 8` rejects seats `1` and `10` without separate comparisons
* Seat `1` becomes `-1`, which is greater than `7` after conversion to `uint`
* Seat `10` becomes `8`, which is not less than `8`
* `2 * n - len(rows)` never iterates empty rows
* The three-mask check is the only extra decrement

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Aisle seats `1` and `10` only, which still allow two families
* A single reserved seat in `4..5`, which leaves the right block free
* Left and right blocked with the middle block free
* All three blocks blocked

## Edge Cases

Important cases to consider:

* `n = 1`
* Reservations only on seats `1` and `10`
* A single reserved seat in the overlap `4..7`
* Both side blocks blocked while the middle remains free
* Left, middle, and right all intersecting a reservation
* `n` far larger than the number of reserved rows

## Notes

* Sorting reserved seats or iterating `1..n` would either be slower or time out for `n = 10^9`.
* Dynamic programming and concurrency add nothing: each row is decided independently in constant time.
* The greedy choice is forced: two side families beat one middle family, and a mapped row never has both sides free.
