---
id: 3620
title: "Network Recovery Pathways"
difficulty: "Hard"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/network-recovery-pathways/"
status: "Solved"
language: "Go"
topics:
  - "Graph"
  - "Binary Search"
  - "Dynamic Programming"
  - "Topological Sort"
go_concepts:
  - "CSR-style adjacency representation"
  - "Binary search over unique edge costs"
  - "Topological shortest-path DP"
  - "int64 for accumulated path costs"
  - "Generation-stamped reachability"
  - "Table-driven and randomized differential tests"
tags:
  - leetcode
  - go
  - graph
  - binary-search
  - topological-sort
  - dynamic-programming
---

# 3620. Network Recovery Pathways

## Problem Link

LeetCode: `https://leetcode.com/problems/network-recovery-pathways/`

## Difficulty

Hard

## Problem Topics

* Graph
* Binary Search
* Dynamic Programming
* Topological Sort

## What to Know Before Solving

General concepts:

* The graph is a directed acyclic graph (DAG)
* A valid path must use only online intermediate nodes
* Total edge cost on the path must be at most `k`
* Each valid path has a **score** equal to the minimum edge cost on that path
* The answer is the maximum score among all valid paths
* Fixing a minimum edge threshold turns the problem into a shortest-path question

Go concepts:

* Building a CSR-style adjacency list with one contiguous edge array
* Binary searching over sorted unique edge costs
* Running shortest-path DP in topological order
* Using `int64` for accumulated costs because `k` can be up to `5 * 10^13`
* Reusing slices with a generation counter instead of clearing them every check
* Writing brute-force and randomized differential tests in the same package

## Problem Description

You are given a directed acyclic graph of `n` nodes numbered from `0` to `n - 1`, represented by `edges`, where `edges[i] = [ui, vi, costi]` is a one-way edge from `ui` to `vi` with recovery cost `costi`.

Some nodes may be offline. Node `0` and node `n - 1` are always online.

A path from `0` to `n - 1` is valid if:

* every intermediate node on the path is online
* the total recovery cost of all edges on the path is at most `k`

For each valid path, define its score as the minimum edge cost along that path.

Return the maximum path score among all valid paths. If no valid path exists, return `-1`.

## Function Signature

Expected LeetCode function signature:

```go
func findMaxPathScore(edges [][]int, online []bool, k int64) int {

}
```

## Examples

### Example 1

Input:

```text
edges = [[0,1,5],[1,3,10],[0,2,3],[2,3,4]], online = [true,true,true,true], k = 10
```

Output:

```text
3
```

Explanation:

```text
Path 0 -> 1 -> 3 has total cost 15, so it is invalid.

Path 0 -> 2 -> 3 has total cost 7 <= k, so it is valid.
Its score is min(3, 4) = 3.

The answer is 3.
```

### Example 2

Input:

```text
edges = [[0,1,7],[1,4,5],[0,2,6],[2,3,6],[3,4,2],[2,4,6]], online = [true,true,true,false,true], k = 12
```

Output:

```text
6
```

Explanation:

```text
Node 3 is offline, so any path through it is invalid.

Valid paths:
- 0 -> 1 -> 4 with score 5
- 0 -> 2 -> 4 with score 6

The answer is 6.
```

## Constraints

```text
n == online.length
2 <= n <= 5 * 10^4
0 <= m == edges.length <= min(10^5, n * (n - 1) / 2)
0 <= ui, vi < n
ui != vi
0 <= costi <= 10^9
0 <= k <= 5 * 10^13
online[0] and online[n - 1] are true
The graph is a DAG
```

## Approach

Each valid path has two numbers:

* **Total cost** — sum of all edge costs on the path (must be `<= k`)
* **Score** — minimum edge cost on the path (what we want to maximize)

Trying every `(totalCost, bottleneck)` pair is too expensive because many different prefixes can share the same bottleneck but differ in total cost.

The key observation is monotonicity:

1. Fix a candidate threshold `x`.
2. Keep only edges with cost `>= x`.
3. Every path built from those edges has score at least `x`.
4. Ask: can we reach `n - 1` from `0` with total cost at most `k`?

That turns the bottleneck question into a shortest-path check on a filtered DAG.

If threshold `x` works, every smaller threshold also works because the same path still exists. If `x` fails, every larger threshold fails because fewer edges are allowed.

The answer must be the cost of some edge on an optimal path, so binary searching over unique edge costs is enough.

## Algorithm

### Preprocess the graph

1. Drop edges touching offline nodes and edges with cost greater than `k`.
2. Build a CSR adjacency list in one contiguous edge array.
3. Compute a topological order with Kahn's algorithm.
4. Collect unique edge costs and sort them for binary search.

### Feasibility check for threshold `x`

1. Reset reachability with a generation-stamped `seen` array.
2. Set `dist[0] = 0`.
3. Walk nodes in topological order.
4. Relax only edges with cost `>= x`.
5. Skip relaxations that would exceed budget `k`.
6. Return `true` as soon as node `n - 1` is reached.

Among all paths that share the same bottleneck, only the cheapest total cost matters. Nonnegative edge costs make the min-cost DP correct on a DAG.

### Binary search

1. Binary search the sorted unique edge costs.
2. Keep the largest feasible threshold.
3. Return `-1` if no threshold is feasible.

## Why This Works

When threshold `x` is fixed, every used edge has cost at least `x`, so the path score is at least `x`.

Among all qualifying paths to the same node, only the cheapest total cost matters. Edge costs are nonnegative, so a more expensive prefix can never beat a cheaper one when extended by the same future edges.

If threshold `x` is feasible, every smaller threshold is also feasible because the same path still works. If `x` is infeasible, every larger threshold is infeasible because fewer edges are allowed.

Therefore, binary searching the maximum feasible edge-cost threshold returns the answer.

## Complexity Analysis

Let:

* `n` = number of nodes
* `m` = number of edges
* `U` = number of unique candidate edge costs

### Time Complexity

```text
O(m log m + (n + m) log U)
```

Building the graph and topological order takes `O(n + m)`. Sorting candidates takes `O(m log m)`. Each feasibility check takes `O(n + m)`, and there are `O(log U)` checks.

### Space Complexity

```text
O(n + m)
```

The CSR adjacency list, topological order, distance array, and candidate costs all fit in linear space.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `edges` and `online` are not mutated
* accumulated path costs use `int64` because `k` can be up to `5 * 10^13`
* CSR layout stores all edges in one slice for better cache locality
* `seen` reuses the indegree slice with a generation counter instead of clearing every check
* early return when the target is reached during relaxation

```go
package networkrecoverypathways

import "sort"

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	type arc struct {
		to   int32
		cost int32
	}

	n := len(online)
	target := n - 1

	offset := make([]int32, n+1)
	indegree := make([]int32, n)

	candidates := make([]int, 0, len(edges))
	validEdgeCount := 0

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		cost := edge[2]

		if !online[from] ||
			!online[to] ||
			int64(cost) > k {
			continue
		}

		offset[from+1]++
		indegree[to]++
		candidates = append(candidates, cost)
		validEdgeCount++
	}

	if validEdgeCount == 0 {
		return -1
	}

	for node := 1; node <= n; node++ {
		offset[node] += offset[node-1]
	}

	adjacency := make([]arc, validEdgeCount)
	order := make([]int32, n)
	copy(order, offset[:n])

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		cost := edge[2]

		if !online[from] ||
			!online[to] ||
			int64(cost) > k {
			continue
		}

		position := order[from]
		adjacency[position] = arc{to: int32(to), cost: int32(cost)}
		order[from]++
	}

	head := 0
	tail := 0

	for node := 0; node < n; node++ {
		if online[node] && indegree[node] == 0 {
			order[tail] = int32(node)
			tail++
		}
	}

	for head < tail {
		node := int(order[head])
		head++

		for index := int(offset[node]); index < int(offset[node+1]); index++ {
			nextNode := int(adjacency[index].to)
			indegree[nextNode]--

			if indegree[nextNode] == 0 {
				order[tail] = int32(nextNode)
				tail++
			}
		}
	}

	order = order[:tail]

	sort.Ints(candidates)

	uniqueCount := 0
	for _, cost := range candidates {
		if uniqueCount == 0 || candidates[uniqueCount-1] != cost {
			candidates[uniqueCount] = cost
			uniqueCount++
		}
	}
	candidates = candidates[:uniqueCount]

	dist := make([]int64, n)
	seen := indegree
	var generation int32

	feasible := func(threshold int) bool {
		generation++
		seen[0] = generation
		dist[0] = 0

		for _, node32 := range order {
			node := int(node32)
			if seen[node] != generation {
				continue
			}

			currentCost := dist[node]

			for index := int(offset[node]); index < int(offset[node+1]); index++ {
				edge := adjacency[index]
				if int(edge.cost) < threshold {
					continue
				}

				newCost := currentCost + int64(edge.cost)
				if newCost > k {
					continue
				}

				nextNode := int(edge.to)
				if nextNode == target {
					return true
				}

				if seen[nextNode] != generation || newCost < dist[nextNode] {
					seen[nextNode] = generation
					dist[nextNode] = newCost
				}
			}
		}

		return false
	}

	answer := -1
	left := 0
	right := len(candidates) - 1

	for left <= right {
		middle := left + (right-left)/2
		threshold := candidates[middle]

		if feasible(threshold) {
			answer = threshold
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return answer
}
```

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* both LeetCode examples
* zero-cost paths with `k = 0`
* offline intermediate nodes
* duplicate edge costs
* overflow-sensitive totals with `int64`
* graphs where a cheap prefix has a lower bottleneck
* 200 seeded randomized differential tests against brute force

Benchmark (`solution_bench_test.go`):

```bash
go test ./problems/3620-network-recovery-pathways -bench=. -benchmem
```

On a chain graph with `n = 5000`, the solution runs in about `490µs/op` with `6` allocations.

## Edge Cases

Important cases to consider:

* no usable edges
* direct edge cost greater than `k`
* offline node blocking all routes
* zero-cost valid path
* multiple valid paths with different bottlenecks
* path total exactly equal to `k`
* large totals requiring `int64`
* disconnected online components

## Notes

* This is a binary search plus topological DP problem, not a plain shortest-path problem.
* Dijkstra is unnecessary because the filtered graph is a DAG.
* The CSR layout avoids one allocation per node and improves cache locality during feasibility checks.
* Do not use cyclic graphs in tests; topological ordering assumes a DAG.
