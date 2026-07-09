---
id: 3532
title: "Path Existence Queries in a Graph I"
difficulty: "Medium"
level: "Senior"
platform: "LeetCode"
link: "https://leetcode.com/problems/path-existence-queries-in-a-graph-i/"
contest: "Weekly Contest 447"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Hash Table"
  - "Binary Search"
  - "Union Find"
  - "Graph Theory"
go_concepts:
  - "Component labeling on a sorted array"
  - "Adjacent gap detection"
  - "uint32 component identifiers"
  - "Table-driven tests with union-find reference"
tags:
  - leetcode
  - go
  - array
  - graph
  - union-find
  - component-labeling
  - weekly-contest-447
---

# 3532. Path Existence Queries in a Graph I

## Problem Link

LeetCode: `https://leetcode.com/problems/path-existence-queries-in-a-graph-i/`

## Difficulty

Medium

## Problem Topics

* Array
* Hash Table
* Binary Search
* Union Find
* Graph Theory

## What to Know Before Solving

General concepts:

* Graph connectivity with edges based on value differences
* How sorted order simplifies connectivity on a line
* Why only adjacent gaps matter for connected components
* Answering many connectivity queries after preprocessing

Go concepts:

* Building a component array in one pass
* Using `uint32` for compact component labels
* Comparing query endpoints in `O(1)`
* Union-find reference tests for randomized validation

## Problem Description

You are given `n` nodes labeled `0` to `n - 1`, a sorted array `nums`, and an integer `maxDiff`.

An undirected edge exists between nodes `i` and `j` when:

```text
|nums[i] - nums[j]| <= maxDiff
```

For each query `[ui, vi]`, return whether a path exists between the two nodes.

## Function Signature

Expected LeetCode function signature:

```go
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {

}
```

## Examples

### Example 1

Input:

```text
n = 2, nums = [1,3], maxDiff = 1, queries = [[0,0],[0,1]]
```

Output:

```text
[true,false]
```

### Example 2

Input:

```text
n = 4, nums = [2,5,6,8], maxDiff = 2, queries = [[0,1],[0,2],[1,3],[2,3]]
```

Output:

```text
[false,false,true,true]
```

## Constraints

```text
1 <= n == nums.length <= 10^5
0 <= nums[i] <= 10^5
nums is sorted in non-decreasing order.
0 <= maxDiff <= 10^5
1 <= queries.length <= 10^5
queries[i] == [ui, vi]
0 <= ui, vi < n
```

## Approach

Because `nums` is sorted, any path from a smaller index to a larger index must pass through every intermediate node on the line.

If some adjacent pair `nums[i] - nums[i-1]` is greater than `maxDiff`, there is no edge between those neighbors, so the graph splits into separate components.

Label each position with its component id while scanning adjacent gaps once. Two nodes are connected iff they share the same component label.

## Algorithm

1. Initialize `component[0] = 0` and `current = 0`.
2. For `i` from `1` to `n - 1`:
   * If `nums[i] - nums[i-1] > maxDiff`, increment `current`.
   * Set `component[i] = current`.
3. For each query `[u, v]`, answer `component[u] == component[v]`.

## Why This Works

### Sorted-array connectivity

For `i < j`, every path from `i` to `j` must use only intermediate nodes with increasing indices. If any adjacent gap `nums[k] - nums[k-1]` exceeds `maxDiff`, nodes on opposite sides of that gap cannot be connected.

Conversely, if every adjacent gap in a segment is at most `maxDiff`, the segment forms one connected component because consecutive nodes are linked by edges.

### Query answering

After component labeling, connectivity reduces to equality of two integers. Each query is answered in constant time.

### Compared with union-find

Union-find also works, but with sorted values and index-line structure, a single linear scan is simpler and faster to implement.

## Complexity Analysis

Let:

```text
n = len(nums)
q = len(queries)
```

### Time Complexity

```text
O(n + q)
```

Component labeling is `O(n)`. Each query is `O(1)`.

### Space Complexity

```text
O(n)
```

For the `component` array.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Only adjacent differences need to be checked because `nums` is sorted
* `component[0]` stays `0` implicitly
* Self-queries are always `true` because both endpoints share the same label

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Single-node graph
* Fully connected chain
* Multiple disconnected components
* Randomized differential tests against union-find

## Edge Cases

Important cases to consider:

* `n = 1`
* `maxDiff = 0` with duplicate values
* Queries where `u == v`
* Nodes in the same component but not directly adjacent
* Nodes in different components separated by one large gap

## Notes

* This is Part I of the path-existence query family on sorted value graphs.
* Binary search and union-find are alternative approaches, but component labeling is optimal here.
* The graph is implicit; no adjacency list is built.
