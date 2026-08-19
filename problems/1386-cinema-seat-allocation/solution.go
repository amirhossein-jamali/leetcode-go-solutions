package cinemaseatallocation

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := make(map[int]uint8, len(reservedSeats))

	for _, reserved := range reservedSeats {
		seat := reserved[1] - 2

		if uint(seat) < 8 {
			rows[reserved[0]] |= uint8(1) << seat
		}
	}

	result := 2*n - len(rows)

	for _, occupied := range rows {
		if occupied&0x0F != 0 &&
			occupied&0x3C != 0 &&
			occupied&0xF0 != 0 {
			result--
		}
	}

	return result
}
