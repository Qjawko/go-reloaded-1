package student

// ConvertBase convert nbr from baseFrom to baseTo
func ConvertBase(nbr, baseFrom, baseTo string) string {
	return nbrToBase(nbrFromBase(nbr, baseFrom), baseTo)
}

func nbrFromBase(nbr, baseFrom string) int {
	baseFromLen := 0
	for range baseFrom {
		baseFromLen++
	}

	nbr = RemoveLeadingChar(nbr, baseFrom[0])

	newNum := 0

	for _, c := range nbr {
		for i, b := range baseFrom {
			if b == c {
				// check for overflow
				if newNum > (MaxInt64-i)/baseFromLen {
					return 0
				}

				newNum = newNum*baseFromLen + i
			}
		}
	}

	return newNum
}

func nbrToBase(nbr int, baseTo string) string {
	if !isValidBase(baseTo) {
		return "NV"
	}

	baseLen := 0
	for range baseTo {
		baseLen++
	}

	r := ""

	if nbr == 0 {
		return string(baseTo[0])
	}

	for nbr != 0 {
		r = string(baseTo[nbr%baseLen]) + r
		nbr /= baseLen
	}

	return r
}
