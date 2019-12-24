package student

// Split mimics Split
func Split(str, charset string) []string {

	strLen := 0
	for range str {
		strLen++
	}

	charsetLen := 0
	for range charset {
		charsetLen++
	}

	// special case if charset length is 0
	if charsetLen == 0 {
		special := make([]string, strLen)

		for i, c := range str {
			special[i] = string(c)
		}

		return special
	}

	// count words
	wordCount := 1
	for i := 0; i < strLen; i++ {
		flag := false

		if str[i] == charset[0] {
			flag = true

			for j := 1; j < charsetLen; j++ {
				if str[i+j] != charset[j] {
					flag = false
				}
			}
		}

		if flag {
			wordCount++
		}
	}

	// main logic
	r := make([]string, wordCount)
	currentIndex := 0

	for i := 0; i < strLen; i++ {
		flag := false

		if str[i] == charset[0] {
			flag = true

			for j := 0; j < charsetLen; j++ {
				if str[i+j] != charset[j] {
					flag = false
				}
			}
		}

		if flag {
			// -1 accounts for i++ at the end of loop
			i += charsetLen - 1
			currentIndex++
		} else {
			r[currentIndex] += string(str[i])
		}
	}

	return r
}
