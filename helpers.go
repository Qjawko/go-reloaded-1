package student

import "github.com/01-edu/z01"

// Print print string
func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

// PrintNbr print number
func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}

	if n != 0 {
		PrintNbr(n / 10)
		z01.PrintRune(rune(n%10 + 48))
	}
}
