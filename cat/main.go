package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	aLen := 0
	for range args {
		aLen++
	}

	if aLen == 0 {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			fmt.Print(text)
		}
	}

	for _, filename := range args {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Println(err)
		}

		print(string(data))
	}

}

func print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
