package main

import (
	"io/ioutil"
	"os"

	student ".."
)

func main() {
	args := os.Args[1:]

	aLen := 0
	for range args {
		aLen++
	}

	if aLen == 0 {
		const bufSize = 512
		bufOut := make([]byte, bufSize)
		n := 0

		for {
			os.Stdin.Read(bufOut[n : n+1])

			if bufOut[n] == '\n' || n == bufSize {
				student.Print(string(bufOut))
				n = 0
			} else {
				n++
			}
		}
	}

	for _, fileName := range args {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			student.Print(err.Error() + "\n")
			continue
		}
		student.Print(string(bytes))
	}
}
