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

		adjacency[position] = arc{
			to:   int32(to),
			cost: int32(cost),
		}

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
