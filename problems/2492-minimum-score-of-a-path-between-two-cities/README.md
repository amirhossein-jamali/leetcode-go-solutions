---
id: 2492
title: "Minimum Score of a Path Between Two Cities"
difficulty: "Medium"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/"
contest: "Weekly Contest 322"
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
  - "Bit packing with int32"
  - "Bit shifts and masks"
  - "Single-pass edge processing"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - graph
  - union-find
  - disjoint-set-union
  - bit-packing
  - weekly-contest-322
---

# 2492. Minimum Score of a Path Between Two Cities

## Problem Link

LeetCode: `https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/`

## Difficulty

Medium

## Problem Topics

* Depth-First Search
* Breadth-First Search
* Union Find
* Graph Theory

## What to Know Before Solving

General concepts:

* The graph is undirected and may contain multiple connected components
* City `1` and city `n` are guaranteed to be connected
* A path may revisit cities and roads
* The score of a path is the minimum edge weight on that path
* The answer is the minimum possible score among all valid walks from city `1` to city `n`

Go concepts:

* Disjoint Set Union with union by size
* Iterative path halving instead of recursive path compression
* Packing multiple metadata fields into one `int32` value
* Processing edges in one pass without building an adjacency list
* Table-driven tests with a BFS reference implementation

## Problem Description

You are given `n` cities numbered from `1` to `n`.

Each road is represented as `[a, b, distance]` and describes a bidirectional weighted road between cities `a` and `b`.

The graph is not necessarily connected, but city `1` and city `n` are guaranteed to be connected.

The score of a path is the minimum road weight contained in that path.

A path may:

* visit the same city multiple times
* traverse the same road multiple times
* revisit cities `1` and `n`

Return the minimum possible score among all paths from city `1` to city `n`.

## Function Signature

Expected LeetCode function signature:

```go
func minScore(n int, roads [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
n = 4
roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
```

Output:

```text
5
```

Explanation:

```text
The path 1 -> 2 -> 4 uses edges with weights 9 and 5.
The score is min(9, 5) = 5.
```

### Example 2

Input:

```text
n = 4
roads = [[1,2,2],[1,3,4],[3,4,7]]
```

Output:

```text
2
```

Explanation:

```text
The direct edge 1 -> 4 does not exist.

A valid walk is:

1 -> 2 -> 1 -> 3 -> 4

This uses edges with weights 2, 2, 4, and 7.
The score is min(2, 2, 4, 7) = 2.

This example shows that the answer edge does not need to belong to a simple path from 1 to n.
```

## Constraints

```text
2 <= n <= 100000
1 <= roads.length <= 100000
roads[i].length == 3
1 <= ai, bi <= n
ai != bi
1 <= distancei <= 10000
There are no repeated edges.
There is at least one path between city 1 and city n.
```

## Core Observation

Because repeated cities and roads are allowed, any edge inside the connected component containing city `1` can be included in some valid walk from city `1` to city `n`.

For an arbitrary edge `(u, v, w)` in that component:

1. There is a path from city `1` to `u`.
2. Traverse the edge `u -> v`.
3. There is a path from `v` to city `n`.

These path segments may overlap or repeat edges, which is allowed.

Therefore, every edge in the connected component containing cities `1` and `n` can participate in a valid walk from `1` to `n`.

Consequently:

```text
answer = minimum edge weight in the connected component containing city 1
```

The algorithm does not need to construct the actual path.

## Approach

1. Process all roads with a Disjoint Set Union structure.
2. Track, for each component, the minimum edge weight seen so far.
3. After all roads are processed, read the minimum edge stored at the root of city `1`.
4. Return that value.

A BFS or DFS solution also works conceptually:

1. Build an adjacency list.
2. Traverse from city `1`.
3. Scan all roads whose endpoints lie in that component.
4. Return the smallest edge weight.

The packed DSU approach avoids the adjacency list, visited state, and queue. It uses one `[]int32` array and processes roads in a single pass.

## Packed DSU Representation

The DSU stores parent pointers and root metadata in one `[]int32` array.

### Non-root node

```text
dsu[node] >= 0
```

The value stores the index of its parent:

```text
dsu[node] = parentIndex
```

### Root node

```text
dsu[root] < 0
```

The negative value stores packed component metadata:

```text
meta = -dsu[root]
componentSize = meta >> metaShift
minimumEdge   = meta & metaMask
```

The packed format is:

```text
meta = (componentSize << 14) | minimumEdge
```

### Why 14 bits are enough for edge weight

```text
maximum distance = 10000
2^14 = 16384
10000 < 16384
```

The sentinel is:

```text
infEdge = 10001
```

which is strictly larger than every valid edge weight.

### Why the packed value fits in int32

```text
maximum component size = 100000

(100000 << 14) + 10001
= 1,638,410,001

math.MaxInt32 = 2,147,483,647
```

So the packed metadata cannot overflow `int32` under the problem constraints.

## Algorithm

1. Initialize every city as a singleton root with `minimumEdge = infEdge`.
2. For each road `[a, b, weight]`:
   * convert cities to zero-based indices
   * find the roots of `a` and `b`
   * update component metadata
3. Return the minimum edge stored at `findRoot(city 1)`.

### Case 1: Both cities already belong to the same component

Only update that component's minimum edge:

```text
componentMin = min(componentMin, weight)
```

This case is necessary because an edge connecting two already-connected cities may still be the minimum edge in the component.

Example:

```text
roads = [[1,2,9],[2,3,8],[1,3,2],[3,4,7]]
```

The edge `[1,3,2]` appears after cities `1`, `2`, and `3` are already connected. That edge lowers the component minimum from `8` to `2`.

### Case 2: The cities belong to different components

Use union by size and combine metadata:

```text
newSize = sizeA + sizeB
newMinimum = min(minimumA, minimumB, currentRoadWeight)
```

Attach the smaller component to the larger component.

### Find operation

`findRoot` uses iterative path halving.

While moving toward the root, a node is redirected to its grandparent whenever possible:

```text
if dsu[parent] >= 0 {
    dsu[node] = dsu[parent]
}
```

This shortens future lookup paths without recursion and gives nearly constant amortized DSU operations.

Do not call this full recursive path compression. The correct term is **iterative path halving**.

## Why This Works

### Lemma 1

Every edge in the connected component containing city `1` can be included in a valid walk from city `1` to city `n`.

**Proof sketch.** Let `(u, v, w)` be any edge in the component. Because the component is connected:

* there exists a walk from city `1` to `u`
* there exists a walk from `v` to city `n`

Concatenate:

```text
walk(1 -> u) + edge(u -> v) + walk(v -> n)
```

Revisiting cities and roads is allowed, so this is a valid walk from `1` to `n` that includes edge `(u, v, w)`.

### Lemma 2

The minimum possible path score equals the minimum edge weight in that connected component.

**Proof sketch.**

* **Upper bound:** Let `w*` be the minimum edge weight in the component. By Lemma 1, some valid walk includes that edge, so the answer is at most `w*`.
* **Lower bound:** Every valid walk from `1` to `n` stays inside the component and uses only edges from that component. The score of any walk is the minimum edge on that walk, which cannot be smaller than the smallest edge in the component.

Therefore the answer equals the component minimum.

### Lemma 3

After all roads have been processed, DSU groups exactly the cities belonging to each connected component.

**Proof sketch.** DSU starts with each city in its own set. Every road performs a union of its two endpoint sets. Two cities end up in the same DSU set if and only if they are connected by some sequence of processed roads. That is exactly the definition of a connected component in an undirected graph.

### Lemma 4

For every DSU root, its packed minimum-edge field equals the minimum weight among all processed edges in that component.

**Proof sketch.** Induction over processed roads.

* **Base case:** A singleton root has no edges yet, so its minimum is `infEdge`.
* **Internal edge case:** If both endpoints already share a root, the root minimum becomes `min(oldMinimum, weight)`.
* **Merge case:** If two different components merge, the new minimum becomes `min(minimumA, minimumB, weight)`.

Both cases preserve the invariant that the root stores the true component minimum.

### Theorem

The minimum-edge metadata stored at `findRoot(city 1)` is exactly the answer.

**Proof.** By Lemma 3, `findRoot(city 1)` is the root of the component containing city `1`. By Lemma 4, its stored minimum equals the smallest edge weight in that component. By Lemma 2, that value is exactly the answer.

## Complexity Analysis

Let:

* `n` = number of cities
* `m` = number of roads

### Time Complexity

```text
O(n + m * alpha(n))
```

Explanation:

* initializing the DSU costs `O(n)`
* each road performs a constant number of find/union operations
* union by size and path halving make each DSU operation amortized `O(alpha(n))`

Here `alpha(n)` is the inverse Ackermann function. It grows so slowly that it is effectively constant for all practical input sizes.

### Auxiliary Space Complexity

```text
O(n)
```

More precisely, the DSU uses one `int32` value per city.

The backing array therefore stores approximately:

```text
4n bytes
```

excluding:

* the input roads
* the Go slice header
* runtime overhead

For `n = 100000`, the DSU backing array is approximately:

```text
400000 bytes
≈ 0.38 MiB
```

Do not claim that the total LeetCode memory usage is only `0.38 MiB`, because the input representation and Go runtime also consume memory.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* cities are converted to zero-based indices internally
* non-negative DSU entries are parent indices
* negative DSU entries are packed root metadata
* union by size keeps trees shallow
* `findRoot` uses iterative path halving
* no adjacency list, visited array, queue, or extra DSU arrays are allocated

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* both LeetCode examples
* direct connection with minimum `n = 2`
* an irrelevant disconnected component containing a smaller edge
* a branch minimum that requires revisiting cities
* an internal cycle edge that updates the component minimum
* minimum edge appearing before later component merges
* minimum edge coming from the second component during a merge
* several edges with maximum allowed weight
* 200 seeded randomized differential tests against a BFS reference implementation

## Edge Cases

Important cases to consider:

* `n = 2` with a single edge
* disconnected components with smaller edges outside city `1`'s component
* minimum edge on a branch, not on a simple `1 -> n` path
* internal edge between already-connected cities lowering the component minimum
* merge where the smaller minimum comes from the attached component
* all edges at maximum weight `10000`

## Notes

* BFS is conceptually simpler: build an adjacency list, traverse from city `1`, then scan roads in the component.
* BFS requires an adjacency list, visited state, and a queue.
* This packed DSU avoids the adjacency list and uses only one `int32` array.
* The packed implementation is more memory-efficient but less immediately readable.
* This optimization is safe because it relies directly on the given constraints:
  * edge weights fit in 14 bits
  * component sizes fit in the remaining upper bits without `int32` overflow

## Related Problems

* [2812. Find the Safest Path in a Grid](./../2812-find-the-safest-path-in-a-grid/) — another graph problem that uses DSU, but for connectivity during descending threshold activation rather than component minimum tracking.
