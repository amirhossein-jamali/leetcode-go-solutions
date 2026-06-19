package findthehighestaltitude

func largestAltitude(gain []int) int {
	currentAltitude := 0
	highestAltitude := 0

	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain

		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}
