---
id: 2685
title: "Count the Number of Complete Components"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/count-the-number-of-complete-components/"
contest: "Weekly Contest 345"
status: "Solved"
language: "Go"
topics:
  - "Depth-First Search"
  - "Breadth-First Search"
  - "Union Find"
  - "Graph Theory"
go_concepts:
  - "Disjoint Set Union"
  - "Union by size"
  - "Iterative path halving"
  - "Bit packing with uint32"
  - "Bit shifts and masks"
  - "Fixed-size array on the stack"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - graph
  - union-find
  - complete-graph
  - bit-packing
  - weekly-contest-345
---

# 2685. Count the Number of Complete Components

## Problem Link

LeetCode: `https://leetcode.com/problems/count-the-number-of-complete-components/`

## Difficulty

Medium

## Problem Topics

* Depth-First Search
* Breadth-First Search
* Union Find
* Graph Theory

## What to Know Before Solving

General concepts:

* Connected components in an undirected graph
* A complete graph has an edge between every pair of distinct vertices
* For a component with `v` vertices, a complete graph must contain exactly `v * (v - 1) / 2` edges
* Counting edges per component is enough once connectivity is known
* A single isolated vertex is a complete component with size `1` and `0` edges

Go concepts:

* Disjoint Set Union with union by size
* Iterative path halving instead of recursive path compression
* Packing component size and edge count into one `uint32`
* Using a fixed `[50]uint32` array because `n <= 50`
* Table-driven tests with a BFS reference implementation

## Problem Description

You are given an integer `n` and an undirected graph with vertices numbered from `0` to `n - 1`.

Each entry `edges[i] = [ai, bi]` means there is an undirected edge between vertices `ai` and `bi`.

Return the number of connected components that are complete.

A connected component is a maximal subgraph where every vertex can reach every other vertex inside the subgraph.

A connected component is complete if there is an edge between every pair of its vertices.

## Function Signature

Expected LeetCode function signature:

```go
func countCompleteComponents(n int, edges [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]
```

Output:

```text
3
```

Explanation:

```text
Component {0, 1, 2} has 3 vertices and 3 edges, so it is complete.
Component {3, 4} has 2 vertices and 1 edge, so it is complete.
Vertex 5 is isolated with size 1 and 0 edges, so it is also complete.
```

### Example 2

Input:

```text
n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]
```

Output:

```text
1
```

Explanation:

```text
Component {0, 1, 2} is complete.
Component {3, 4, 5} has 3 vertices but only 2 edges, so it is not complete.
Only one complete component exists.
```

## Constraints

```text
1 <= n <= 50
0 <= edges.length <= n * (n - 1) / 2
edges[i].length == 2
0 <= ai, bi <= n - 1
ai != bi
There are no repeated edges.
```

## Approach

Use Union-Find to merge vertices connected by edges and maintain two values for each component root:

* component size
* number of edges inside the component

After processing all edges, each root represents one connected component. A component is complete iff:

```text
edgeCount == size * (size - 1) / 2
```

Because `n <= 50`, all DSU state can live in a fixed-size `uint32` array with no heap allocation.

## Algorithm

1. Initialize every vertex as a singleton root with `size = 1` and `edgeCount = 0`.
2. For each edge `[a, b]`:
   * Find the roots of `a` and `b` with path halving.
   * If both vertices already belong to the same component, increment that root's edge count and continue.
   * Otherwise, union by size:
     * Attach the smaller component to the larger one.
     * Set the new edge count to `edgesA + edgesB + 1`.
3. Scan all vertices.
4. For every root:
   * Read `size` and `edgeCount`.
   * If `edgeCount == size * (size - 1) / 2`, increment the answer.
5. Return the answer.

## Why This Works

### Connectivity

Union-Find correctly groups vertices into connected components. Every edge either:

* connects two different components, which are merged, or
* lies inside one component, which increases that component's internal edge count.

So after all edges are processed, each root corresponds to exactly one connected component.

### Completeness test

In an undirected simple graph, a connected component with `v` vertices can have at most `v * (v - 1) / 2` edges, and it reaches that maximum only when every pair of vertices is connected.

Therefore:

* if `edgeCount == v * (v - 1) / 2`, the component is complete
* otherwise, at least one pair is missing an edge, so it is not complete

### Singleton components

A vertex with no edges has:

```text
size = 1
edgeCount = 0
0 == 1 * 0 / 2
```

So isolated vertices are counted correctly as complete components.

### Packed metadata

Each root stores:

* bit `31`: root marker
* lower `11` bits: edge count
* next `6` bits: component size

This allows union, find, and final counting in one array without separate maps or adjacency lists.

## Complexity Analysis

Let:

```text
n = number of vertices
m = number of edges
```

### Time Complexity

```text
O(m * alpha(n) + n)
```

Union-Find operations are nearly constant per edge, and the final scan over roots is `O(n)`.

With `n <= 50`, this is effectively linear in the input size.

### Space Complexity

```text
O(1)
```

The DSU state uses a fixed `[50]uint32` array on the stack.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `rootBit` marks packed root nodes; non-root entries store only a parent index
* `findComponent` uses iterative path halving for cache-friendly find operations
* Union by size keeps trees shallow and simplifies merging metadata
* When an edge connects two vertices already in the same component, only the edge count increases
* `edgeMask` and `sizeMask` decode packed root metadata during the final completeness check

### Bit layout of a root node

```text
bit 31          : root flag
bits 11..16     : component size (6 bits, enough for n <= 50)
bits 0..10      : edge count inside the component (11 bits)
```

Example for a singleton root:

```text
size = 1, edges = 0
root = 1<<31 | 1<<11 | 0
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* Both LeetCode examples
* Single isolated vertex
* Multiple isolated vertices
* Complete graph on 3 and 4 vertices
* A non-complete path component mixed with a complete singleton
* Two disconnected complete pairs
* A star graph, which is connected but not complete
* Randomized differential tests against a BFS reference

## Edge Cases

Important cases to consider:

* `n = 1` with no edges
* Several isolated vertices with no edges at all
* A size-2 component with exactly one edge
* A connected component that is not complete, such as a path or star
* A full clique on the maximum useful size for this problem
* Components where `edgeCount` is close to but less than `v * (v - 1) / 2`

## Notes

* DFS or BFS also works, but you still need to count edges per component afterward.
* Union-Find is a strong fit because connectivity and per-component edge totals are updated in one pass.
* The packed `uint32` representation is especially convenient here because `n <= 50` makes fixed-width metadata safe.
* A common mistake is checking only connectivity and forgetting the exact complete-graph edge formula.
* Another common mistake is counting undirected edges twice when building an adjacency-list-based solution.

## Completion Checklist

* YAML frontmatter is filled
* Problem link, difficulty, topics, and contest are filled
* Approach explains Union-Find with packed metadata
* Function signature matches LeetCode
* Examples and constraints are included
* Time and space complexity are included
* `solution.go` uses the expected LeetCode function signature
* `solution_test.go` covers examples and edge cases
* `go test ./...` passes
