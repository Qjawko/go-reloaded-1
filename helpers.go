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

	r := ""

	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		Print("-")
		for n != 0 {
			r = string(rune((n%10*-1)+48)) + r
			n /= 10
		}
	} else if n > 0 {
		for n != 0 {
			r = string(rune((n%10)+48)) + r
			n /= 10
		}
	}

	Print(r)
}
