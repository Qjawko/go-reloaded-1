package student

// MaxInt64 max int64 value
var MaxInt64 int = 9223372036854775807

// MinInt64 min int64 value
var MinInt64 int = -9223372036854775808

// Atoi basic atoi
func Atoi(s string) int {
	result := 0

	if !isValidNumber(s) {
		return 0
	}

	flag := false

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			flag = true
		}
		s = s[1:]
	}

	s = RemoveLeadingZeroes(s)

	if s == "0" {
		return 0
	}

	for _, c := range s {
		x := int(c - '0')

		if flag {
			// negative case
			// check for overflow
			if result < (MinInt64+x)/10 {
				return 0
			}

			result = result*10 - x
			continue
		}

		// positive case
		// check for overflow
		if result > (MaxInt64-x)/10 {
			return 0
		}
		result = result*10 + x
	}

	return result
}

// RemoveLeadingZeroes from string
func RemoveLeadingZeroes(s string) string {
	len := 0
	for range s {
		len++
	}

	if len < 2 {
		return s
	}

	if s[0] == '0' {
		return RemoveLeadingZeroes(s[1:])
	}

	return s
}

func isValidNumber(s string) bool {
	for i, c := range s {

		if c == '+' || c == '-' && i == 0 {
			continue
		}

		if c < '0' || c > '9' {
			return false
		}

	}

	return true
}
