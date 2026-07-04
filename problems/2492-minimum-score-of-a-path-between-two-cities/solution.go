package minimumscoreofapathbetweentwocities

const (
	metaShift = 14                          // upper bits store component size
	metaMask  = int32((1 << metaShift) - 1) // lower 14 bits store minimum edge weight
	infEdge   = int32(10_001)
)

// findRoot returns the representative of node's set using iterative path halving.
// Non-negative dsu entries are parent indices; negative entries are packed root metadata.
func findRoot(dsu []int32, node int32) int32 {
	for dsu[node] >= 0 {
		parent := dsu[node]

		if dsu[parent] >= 0 {
			dsu[node] = dsu[parent]
		}

		node = dsu[node]
	}

	return node
}

func minScore(n int, roads [][]int) int {
	dsu := make([]int32, n)

	// Each root starts as a singleton with no edges seen yet.
	initialRoot := -((int32(1) << metaShift) | infEdge)

	for city := range dsu {
		dsu[city] = initialRoot
	}

	for _, road := range roads {
		rootA := findRoot(dsu, int32(road[0]-1))
		rootB := findRoot(dsu, int32(road[1]-1))
		weight := int32(road[2])

		metaA := -dsu[rootA]

		if rootA == rootB {
			if weight < metaA&metaMask {
				dsu[rootA] = -((metaA &^ metaMask) | weight)
			}
			continue
		}

		metaB := -dsu[rootB]

		if metaA>>metaShift < metaB>>metaShift {
			rootA, rootB = rootB, rootA
			metaA, metaB = metaB, metaA
		}

		minimum := weight

		if minA := metaA & metaMask; minA < minimum {
			minimum = minA
		}

		if minB := metaB & metaMask; minB < minimum {
			minimum = minB
		}

		newSize := (metaA >> metaShift) + (metaB >> metaShift)

		dsu[rootB] = rootA
		dsu[rootA] = -((newSize << metaShift) | minimum)
	}

	root := findRoot(dsu, 0)

	return int((-dsu[root]) & metaMask)
}
