package maximumicecreambars

func maxIceCream(costs []int, coins int) int {
	maxCost := 0

	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	frequency := make([]int, maxCost+1)

	for _, cost := range costs {
		frequency[cost]++
	}

	bought := 0

	for price := 1; price <= maxCost && coins >= price; price++ {
		if frequency[price] == 0 {
			continue
		}

		count := coins / price

		if count > frequency[price] {
			count = frequency[price]
		}

		bought += count
		coins -= count * price

		if count < frequency[price] {
			break
		}
	}

	return bought
}
