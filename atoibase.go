package student

// AtoiBase s to base
func AtoiBase(s string, base string) int {

	if !isValidBase2(base) {
		return 0
	}

	baseLen := 0
	for range base {
		baseLen++
	}

	sLen := 0
	for range s {
		sLen++
	}

	r := 0
	for i, c := range s {
		for ib, b := range base {
			if c == b {
				r += ib * recursivePower2(baseLen, sLen-1-i)
			}
		}
	}

	return r
}

func recursivePower2(nb int, power int) int {
	if power < 0 || nb == 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	return nb * recursivePower2(nb, power-1)
}

// checks if base is valid
func isValidBase2(base string) bool {
	len := 0
	for range base {
		len++
	}

	if len < 2 {
		return false
	}

	for i := 0; i < len-1; i++ {
		c := base[i]

		for j := i + 1; j < len; j++ {
			c2 := base[j]

			if c2 == c || c2 == '+' || c2 == '-' || c == '+' || c == '-' {
				return false
			}
		}
	}

	return true
}
