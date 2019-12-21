package student

// Compare q
func Compare(s1, s2 string) int {

	if s1 > s2 {
		return 1
	}

	if s2 > s1 {
		return -1
	}

	return 0
}
