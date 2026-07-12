package countthenumberofcompletecomponents

const (
	rootBit  uint32 = 1 << 31
	edgeBits        = 11
	edgeMask uint32 = (1 << edgeBits) - 1
	sizeMask uint32 = 63
)

func makeRoot(size, edges uint32) uint32 {
	return rootBit | size<<edgeBits | edges
}

func componentSize(value uint32) uint32 {
	return (value >> edgeBits) & sizeMask
}

func findComponent(state *[50]uint32, node int) int {
	for state[node]&rootBit == 0 {
		parent := int(state[node])

		if state[parent]&rootBit == 0 {
			state[node] = state[parent]
		}

		node = int(state[node])
	}

	return node
}

func countCompleteComponents(n int, edges [][]int) int {
	var state [50]uint32

	for node := 0; node < n; node++ {
		state[node] = makeRoot(1, 0)
	}

	for _, edge := range edges {
		rootA := findComponent(&state, edge[0])
		rootB := findComponent(&state, edge[1])

		if rootA == rootB {
			state[rootA]++
			continue
		}

		sizeA := componentSize(state[rootA])
		sizeB := componentSize(state[rootB])

		if sizeA < sizeB {
			rootA, rootB = rootB, rootA
			sizeA, sizeB = sizeB, sizeA
		}

		edgeCount := (state[rootA] & edgeMask) +
			(state[rootB] & edgeMask) + 1

		state[rootB] = uint32(rootA)
		state[rootA] = makeRoot(sizeA+sizeB, edgeCount)
	}

	answer := 0

	for node := 0; node < n; node++ {
		if state[node]&rootBit == 0 {
			continue
		}

		size := componentSize(state[node])
		edgeCount := state[node] & edgeMask

		if edgeCount == size*(size-1)/2 {
			answer++
		}
	}

	return answer
}
