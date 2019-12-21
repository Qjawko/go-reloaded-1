package student

// ConvertBase convert nbr from baseFrom to baseTo
func ConvertBase(nbr, baseFrom, baseTo string) string {

	integer := nbrFromBase(nbr, baseFrom)

	newNbr := nbrToBase(integer, baseTo)

	return newNbr
}

func nbrFromBase(nbr, baseFrom string) int {
	// calculate length of base
	baseFromLen := 0
	for range baseFrom {
		baseFromLen++
	}

	sign := 1

	if nbr[0] == '-' || nbr[0] == '+' {
		if nbr[0] == '-' {
			sign = -1
		}

		nbr = nbr[1:]
	}

	newNum := 0

	for _, c := range nbr {
		for i, b := range baseFrom {
			if b == c {
				newNum = i + newNum*baseFromLen
			}
		}
	}

	return sign * newNum
}

func nbrToBase(nbr int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}

	// calculate length of base
	baseLen := 0
	for range base {
		baseLen++
	}

	r := ""
	sign := ""

	if nbr == 0 {
		return string(base[0])
	}

	if nbr < 0 {
		sign = "-"
		nbr *= -1
	}

	for nbr != 0 {
		r = string(base[nbr%baseLen]) + r
		nbr /= baseLen
	}

	return sign + r
}
