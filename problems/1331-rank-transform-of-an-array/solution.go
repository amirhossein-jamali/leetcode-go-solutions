package ranktransformofanarray

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	storage := make([]uint32, 2*n)
	src := storage[:n]
	dst := storage[n:]

	for i := range src {
		src[i] = uint32(i)
	}

	for shift := uint(0); shift < 32; shift += 8 {
		var count [256]int

		for _, index := range src {
			key := uint32(int32(arr[index])) ^ uint32(1<<31)
			bucket := byte(key >> shift)
			count[bucket]++
		}

		position := 0
		for i := range count {
			frequency := count[i]
			count[i] = position
			position += frequency
		}

		for _, index := range src {
			key := uint32(int32(arr[index])) ^ uint32(1<<31)
			bucket := byte(key >> shift)

			dst[count[bucket]] = index
			count[bucket]++
		}

		src, dst = dst, src
	}

	rank := 0
	var previousKey uint32

	for i, index := range src {
		key := uint32(int32(arr[index])) ^ uint32(1<<31)

		if i == 0 || key != previousKey {
			rank++
			previousKey = key
		}

		arr[index] = rank
	}

	return arr
}
