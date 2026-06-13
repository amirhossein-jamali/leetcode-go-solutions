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

| #    | Problem               | Difficulty | Solution                                      |
| ---- | --------------------- | ---------- | --------------------------------------------- |
| 3838 | Weighted Word Mapping | Easy       | [Go](./problems/3838-weighted-word-mapping/) |

## How to Run Tests

Run all tests:

```bash
go test ./...
```

Run tests for a specific problem:

```bash
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
