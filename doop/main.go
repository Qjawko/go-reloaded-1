package main

import (
	"os"

	student ".."
)

/// @TODO add overflow check
func main() {
	args := os.Args[1:]

	argsLen := 0
	for range args {
		argsLen++
	}

	if argsLen != 3 {
		student.Print("0\n")
		return
	}

	n1 := student.Atoi(args[0])
	n2 := student.Atoi(args[2])
	sign := args[1]

	// check if atoi returns 0
	if (n1 == 0 && args[0] != "0") || (n2 == 0 && args[2] != "0") {
		student.Print("0\n")
		return
	}

	switch sign {
	// plus
	case "+":

		student.PrintNbr(n1 + n2)
		student.Print("\n")

	// minus
	case "-":
		student.PrintNbr(n1 - n2)
		student.Print("\n")

	// multiply
	case "*":
		student.PrintNbr(n1 * n2)
		student.Print("\n")

	// divide
	case "/":
		if n2 != 0 {
			student.PrintNbr(n1 / n2)
			student.Print("\n")
		} else {
			student.Print("No division by 0\n")
		}

	// modulo
	case "%":
		if n2 != 0 {
			student.PrintNbr(n1 % n2)
			student.Print("\n")
		} else {
			student.Print("No modulo by 0\n")
		}

	default:
		student.Print("0\n")
		return
	}

}
