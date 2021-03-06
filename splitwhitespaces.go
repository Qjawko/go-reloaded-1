package student

// SplitWhiteSpaces split() with whitespaces
func SplitWhiteSpaces(str string) []string {
	// remove spaces in front and back
	str = trimWhitespaces(str)

	// remove repeating spaces
	str = leaveOneSpace(str)

	// main logic
	// count words
	wordCount := 1
	for _, c := range str {
		if c == ' ' || c == '\t' || c == '\n' {
			wordCount++
		}
	}

	// make slice with length of words
	result := make([]string, wordCount)
	currentIndex := 0

	for _, c := range str {
		if c != ' ' && c != '\t' && c != '\n' {
			result[currentIndex] += string(c)
		} else {
			currentIndex++
		}
	}

	return result
}

func trimWhitespaces(s string) string {
	sLen := 0
	for range s {
		sLen++
	}

	if sLen == 0 {
		return s
	}

	if s[0] == ' ' || s[0] == '\t' || s[0] == '\n' {
		return trimWhitespaces(s[1:])
	}

	if s[sLen-1] == ' ' || s[sLen-1] == '\t' || s[sLen-1] == '\n' {
		return trimWhitespaces(s[:sLen-1])
	}

	return s
}

func leaveOneSpace(s string) string {
	r := ""

	for i, c := range s {
		if (c != ' ' && c != '\t' && c != '\n') || ((c == ' ' || c == '\t' || c == '\n') && (i == 0 || (s[i-1] != ' ' && s[i-1] != '\t' && s[i-1] != '\n'))) {
			r += string(c)
		}
	}

	return r
}
