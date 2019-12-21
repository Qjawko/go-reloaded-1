package student

// PrintNbrBase prints number nbr in base string
func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		print("NV")
		return
	}

	// calculate length of base
	baseLen := 0
	for range base {
		baseLen++
	}

	r := ""
	sign := ""

	if nbr == 0 {
		print(string(base[0]))
	}

	if nbr < 0 {
		sign = "-"
		nbr *= -1
	}

	for nbr != 0 {
		r = string(base[nbr%baseLen]) + r
		nbr /= baseLen
	}

	print(sign + r)
}

// checks if base is valid
func isValidBase(base string) bool {
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
