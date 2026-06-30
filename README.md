# LeetCode Go Solutions

This repository contains my solutions for LeetCode problems written in Go.

The main goal of this repository is to practice problem-solving, improve algorithmic thinking, and keep a clean record of my progress.

## Language

Go

## Repository Structure

```text
leetcode-go-solutions/
├── README.md
├── go.mod
├── .gitignore
├── docs/
│   └── problem-template.md
└── problems/
    ├── 125-valid-palindrome/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1189-maximum-number-of-balloons/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1344-angle-between-hands-of-a-clock/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1358-number-of-substrings-containing-all-three-characters/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1732-find-the-highest-altitude/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1833-maximum-ice-cream-bars/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1840-maximum-building-height/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1846-maximum-element-after-decreasing-and-rearranging/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 1967-number-of-strings-that-appear-as-substrings-in-word/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 2095-delete-the-middle-node-of-a-linked-list/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 2130-maximum-twin-sum-of-a-linked-list/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 2287-rearrange-characters-to-make-target-string/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3020-find-the-maximum-number-of-elements-in-subset/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3612-process-string-with-special-operations-i/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3614-process-string-with-special-operations-ii/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3699-number-of-zigzag-arrays-i/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3700-number-of-zigzag-arrays-ii/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3737-count-subarrays-with-majority-element-i/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    ├── 3739-count-subarrays-with-majority-element-ii/
    │   ├── README.md
    │   ├── solution.go
    │   └── solution_test.go
    └── 3838-weighted-word-mapping/
        ├── README.md
        ├── solution.go
        └── solution_test.go
```

Each problem has its own folder inside the `problems` directory.

The folder name follows this format:

```text
problem-number-problem-name
```

Example:

```text
0001-two-sum
0020-valid-parentheses
0121-best-time-to-buy-and-sell-stock
```

## Problem Folder Structure

Each problem folder contains:

```text
README.md
solution.go
solution_test.go
```

### `README.md`

Contains the problem explanation, approach, algorithm idea, and complexity analysis.

Problem README files can also include Obsidian-friendly YAML frontmatter for metadata such as difficulty, topics, contest, and Go concepts.

### `solution.go`

Contains the Go implementation of the solution.

### `solution_test.go`

Contains local test cases for the solution.

## Solved Problems

| #    | Problem                                   | Difficulty | Solution                                                      |
| ---- | ----------------------------------------- | ---------- | ------------------------------------------------------------- |
| 125  | Valid Palindrome                          | Easy       | [Go](./problems/125-valid-palindrome/)                        |
| 1189 | Maximum Number of Balloons                | Easy       | [Go](./problems/1189-maximum-number-of-balloons/)             |
| 1344 | Angle Between Hands of a Clock            | Medium     | [Go](./problems/1344-angle-between-hands-of-a-clock/)         |
| 1358 | Number of Substrings Containing All Three Characters | Medium | [Go](./problems/1358-number-of-substrings-containing-all-three-characters/) |
| 1732 | Find the Highest Altitude                 | Easy       | [Go](./problems/1732-find-the-highest-altitude/)              |
| 1833 | Maximum Ice Cream Bars                    | Medium     | [Go](./problems/1833-maximum-ice-cream-bars/)                 |
| 1840 | Maximum Building Height                   | Hard       | [Go](./problems/1840-maximum-building-height/)                |
| 1846 | Maximum Element After Decreasing and Rearranging | Medium | [Go](./problems/1846-maximum-element-after-decreasing-and-rearranging/) |
| 1967 | Number of Strings That Appear as Substrings in Word | Easy | [Go](./problems/1967-number-of-strings-that-appear-as-substrings-in-word/) |
| 2095 | Delete the Middle Node of a Linked List   | Medium     | [Go](./problems/2095-delete-the-middle-node-of-a-linked-list/) |
| 2130 | Maximum Twin Sum of a Linked List         | Medium     | [Go](./problems/2130-maximum-twin-sum-of-a-linked-list/)      |
| 2287 | Rearrange Characters to Make Target String | Easy       | [Go](./problems/2287-rearrange-characters-to-make-target-string/) |
| 3020 | Find the Maximum Number of Elements in Subset | Medium | [Go](./problems/3020-find-the-maximum-number-of-elements-in-subset/) |
| 3612 | Process String with Special Operations I  | Medium     | [Go](./problems/3612-process-string-with-special-operations-i/) |
| 3614 | Process String with Special Operations II | Hard       | [Go](./problems/3614-process-string-with-special-operations-ii/) |
| 3699 | Number of ZigZag Arrays I                 | Hard       | [Go](./problems/3699-number-of-zigzag-arrays-i/)              |
| 3700 | Number of ZigZag Arrays II                | Hard       | [Go](./problems/3700-number-of-zigzag-arrays-ii/)             |
| 3737 | Count Subarrays With Majority Element I   | Medium     | [Go](./problems/3737-count-subarrays-with-majority-element-i/) |
| 3739 | Count Subarrays With Majority Element II  | Hard       | [Go](./problems/3739-count-subarrays-with-majority-element-ii/) |
| 3838 | Weighted Word Mapping                     | Easy       | [Go](./problems/3838-weighted-word-mapping/)                  |

## How to Run Tests

Run all tests:

```bash
go test ./...
```

Run tests for one problem:

```bash
go test ./problems/125-valid-palindrome
go test ./problems/1189-maximum-number-of-balloons
go test ./problems/1344-angle-between-hands-of-a-clock
go test ./problems/1358-number-of-substrings-containing-all-three-characters
go test ./problems/1732-find-the-highest-altitude
go test ./problems/1833-maximum-ice-cream-bars
go test ./problems/1840-maximum-building-height
go test ./problems/1846-maximum-element-after-decreasing-and-rearranging
go test ./problems/1967-number-of-strings-that-appear-as-substrings-in-word
go test ./problems/2095-delete-the-middle-node-of-a-linked-list
go test ./problems/2130-maximum-twin-sum-of-a-linked-list
go test ./problems/2287-rearrange-characters-to-make-target-string
go test ./problems/3020-find-the-maximum-number-of-elements-in-subset
go test ./problems/3612-process-string-with-special-operations-i
go test ./problems/3614-process-string-with-special-operations-ii
go test ./problems/3699-number-of-zigzag-arrays-i
go test ./problems/3700-number-of-zigzag-arrays-ii
go test ./problems/3737-count-subarrays-with-majority-element-i
go test ./problems/3739-count-subarrays-with-majority-element-ii
go test ./problems/3838-weighted-word-mapping
```

## Goals

The goals of this repository are:

* Practice data structures and algorithms
* Improve Go programming skills
* Learn how to analyze time and space complexity
* Write clean and readable solutions
* Keep a consistent record of solved problems

## Common Topics

This repository may include problems related to:

* Arrays
* Strings
* Hash Maps
* Two Pointers
* Sliding Window
* Stack
* Queue
* Linked List
* Trees
* Graphs
* Recursion
* Backtracking
* Dynamic Programming
* Binary Search
* Sorting

## Notes

The focus is not only on getting accepted by LeetCode.

The main focus is understanding why each solution works, what trade-offs it has, and how efficient it is.

Some problems may include more than one approach, such as:

* Brute force
* Optimized solution
* Alternative solution

## Progress

I will update this repository as I solve more problems.
