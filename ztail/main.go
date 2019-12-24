package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	aLen := 0
	for range args {
		aLen++
	}

	maxLen := uint(18446744073709551615)
	flag := false

	for i, arg := range args {
		if arg == "-c" {
			if i >= aLen-1 {
				fmt.Printf("%s", "tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
				return
			}

			msg := ""
			maxLen, flag, msg = atoi(args[i+1])

			if maxLen == 0 && !flag {
				return
			}

			if msg != "" {
				fmt.Printf("%v", msg)
			}

			args = append(args[:i], args[i+2:]...)
		}
	}

	aLen = 0
	for range args {
		aLen++
	}

	for i, fileName := range args {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			// fmt.Printf("%v\n", err.Error())
			if err.Error() == "read "+fileName+": The handle is invalid." {
				fmt.Printf("%v\n", "tail: error reading '"+fileName+"': Is a directory")
			} else {
				fmt.Printf("%v\n", "tail: cannot open '"+fileName+"' for reading: No such file or directory")
			}
			continue
		}

		dLen := uint(0)
		for range data {
			dLen++
		}

		result := ""
		lines := -1
		count := uint(0)

		if flag {
			i := maxLen - 1
			if maxLen == 0 {
				i = 0
			}
			for ; i < dLen; i++ {
				result += string(data[i])
			}
		} else {
			for i := dLen - 1; i >= 0 && i < dLen; i-- {
				c := data[i]

				if c == '\n' {
					lines++
				}
				if lines == 10 || count == maxLen {
					break
				}

				result = string(c) + result
				count++
			}
		}

		if i > 0 {
			fmt.Printf("%s", "\n")
		}

		if aLen > 1 {
			fmt.Printf("==> %s <==\n", fileName)
		}

		fmt.Printf("%s", result)
	}
}

func atoi(s string) (uint, bool, string) {
	result := uint(0)
	msg := ""

	if !isValidNumber(s) {
		msg = "tail: invalid number of bytes: ‘" + s + "’\n"
		return 0, false, msg
	}

	flag := false

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '+' {
			flag = true
		}
		s = s[1:]
	}

	s = removeLeadingChar(s, '0')

	if s == "0" {
		return 0, flag, msg
	}

	for _, c := range s {
		x := uint(c - '0')

		// check for overflow
		if result > (18446744073709551615-x)/10 {
			msg = "tail: invalid number of bytes: ‘" + s + "’: Value too large for defined data type\n"
			return 0, flag, msg
		}
		result = result*10 + x
	}

	return result, flag, msg
}

func isValidNumber(s string) bool {
	len := 0
	for range s {
		len++
	}
	for i, c := range s {

		if c == '+' || c == '-' && i == 0 {
			if len > 1 {
				continue
			}
		}

		if c < '0' || c > '9' {
			return false
		}

	}

	return true
}

func removeLeadingChar(s string, char byte) string {
	len := 0
	for range s {
		len++
	}

	if len < 2 {
		return s
	}

	if s[0] == char {
		return removeLeadingChar(s[1:], char)
	}

	return s
}
