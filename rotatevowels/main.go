package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	argsLen := 0
	for range args {
		argsLen++
	}

	if argsLen == 0 {
		print("\n")
		return
	}

	// join into one string
	str := join(args, " ")

	strLen := 0
	for range str {
		strLen++
	}

	// save all vowels
	vowelsStr := ""
	for _, c := range str {
		if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') ||
			(c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U') {
			vowelsStr += string(c)
		}
	}

	vLen := 0
	for range vowelsStr {
		vLen++
	}

	// save all vowels indexi
	indexi := make([]int, vLen)
	curr := 0
	for i, c := range str {
		if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') ||
			(c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U') {
			indexi[curr] = i
			curr++
		}
	}

	// rotate
	strArr := []rune(str)
	for x := range indexi {
		strArr[indexi[vLen-x-1]] = rune(vowelsStr[x])
	}

	// final print
	print(string(strArr) + "\n")
}

func print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func join(arr []string, divider string) string {
	r := ""

	len := 0
	for range arr {
		len++
	}

	for i, item := range arr {

		r += item

		if i != len-1 {
			r += divider
		}
	}
	return r
}
