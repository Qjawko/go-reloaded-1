package student

// AdvancedSortWordArr sorts words with give function
func AdvancedSortWordArr(array []string, f func(a, b string) int) {

	if len(array) < 2 {
		return
	}

	for x := 0; x < len(array)-1; x++ {
		for i := 1; i < len(array)-x; i++ {
			a := array[i-1]
			b := array[i]

			if f(a, b) > 0 {
				array[i], array[i-1] = array[i-1], array[i]
			}
		}
	}
}
