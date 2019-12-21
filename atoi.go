package student

// Atoi basic atoi
func Atoi(s string) int {
	r := 0
	sign := 1

	for i, c := range s {
		if c >= '0' && c <= '9' {
			// prev := r

			r = r*10 + int(c-48)

			// catch overflow
			// if prev != 0 && prev > r {
			// 	return 0
			// }

		} else if i == 0 {
			if c == '-' {
				sign = -1
			} else if c != '+' {
				return 0
			}
		} else {
			return 0
		}
	}

	// fmt.Println(r)

	return r * sign
}
