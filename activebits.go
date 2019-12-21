package student

// ActiveBits return number of active bits
func ActiveBits(n int) uint {
	var count uint
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}

	return count
}
