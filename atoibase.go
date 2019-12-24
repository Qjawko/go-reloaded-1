package student

// AtoiBase s to base
func AtoiBase(nbr string, base string) int {

	// check if base is valid
	if !isValidBase(base) {
		return 0
	}

	// check if base contains nbr characters
	if !checkNumberForBase(nbr, base) {
		return 0
	}

	baseLen := 0
	for range base {
		baseLen++
	}

	sLen := 0
	for range nbr {
		sLen++
	}

	r := 0
	for in, n := range nbr {
		for ib, b := range base {
			if n == b {

				x := ib * RecursivePower(baseLen, sLen-1-in)

				// handle overflow
				if r > MaxInt64-x {
					return 0
				}

				r += x
			}
		}
	}

	return r
}

func checkNumberForBase(nbr, base string) bool {
	for _, n := range nbr {
		flag := false
		for _, b := range base {
			if b == n {
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}

	return true
}
