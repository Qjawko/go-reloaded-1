package main

import (
	"os"

	student ".."
)

func main() {
	// get args
	args := os.Args[1:]

	// get args length
	argsLen := 0
	for range args {
		argsLen++
	}

	// args length should be always 3
	if argsLen != 3 {
		student.Print("\n")
		return
	}

	// trim zeroes
	args[0] = student.RemoveLeadingZeroes(args[0])
	args[2] = student.RemoveLeadingZeroes(args[2])

	n1 := student.Atoi(args[0])
	n2 := student.Atoi(args[2])
	sign := args[1]

	// check if atoi returns 0
	if (n1 == 0 && args[0] != "0") || (n2 == 0 && args[2] != "0") {
		student.Print("0\n")
		return
	}

	// change n2 to positive for plus or minus
	if sign == "+" && n2 < 0 {
		sign = "-"
		n2 *= -1
	}
	if sign == "-" && n2 < 0 {
		sign = "+"
		n2 *= -1
	}

	// handle plus
	if sign == "+" {
		// handle overflow
		if n1 > 0 {
			if n1 > student.MaxInt64-n2 {
				student.Print("0\n")
				return
			}
		}

		student.PrintNbr(n1 + n2)
		student.Print("\n")
		return
	}

	// handle minus
	if sign == "-" {
		// handle overflow
		if n1 < 0 {
			if n1 < student.MinInt64+n2 {
				student.Print("0\n")
				return
			}
		}

		student.PrintNbr(n1 - n2)
		student.Print("\n")
		return
	}

	// handle division
	if sign == "/" {
		// handle division by 0
		if n2 == 0 {
			student.Print("No division by 0\n")
			return
		}

		student.PrintNbr(n1 / n2)
		student.Print("\n")
		return
	}

	// handle modulo
	if sign == "%" {
		// handle modulo by 0
		if n2 == 0 {
			student.Print("No modulo by 0\n")
			return
		}

		student.PrintNbr(n1 % n2)
		student.Print("\n")
		return
	}

	// handle multiplication
	if sign == "*" {
		// handle overflows
		if n1 < 0 && n2 < 0 {
			if n1 < student.MaxInt64/n2 {
				student.Print("0\n")
				return
			}
		}
		if n1 > 0 && n2 > 0 {
			if n1 > student.MaxInt64/n2 {
				student.Print("0\n")
				return
			}
		}
		if n1 < 0 && n2 > 0 {
			if n1 < student.MinInt64/n2 {
				student.Print("0\n")
				return
			}
		}
		if n1 > 0 && n2 < 0 {
			if n2 < student.MinInt64/n1 {
				student.Print("0\n")
				return
			}
		}

		student.PrintNbr(n1 * n2)
		student.Print("\n")
		return
	}

	student.Print("0\n")
}
