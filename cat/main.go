package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if lenArgs(args) == 1 {
		buf := []byte{0}
		if _, err := io.CopyBuffer(os.Stdout, os.Stdin, buf); err != nil {
			for k := range err.Error() {
				z01.PrintRune(rune(err.Error()[k]))
			}
			z01.PrintRune(10)
		}
	} else {
		args = args[1:]
		for i := range args {
			file, err := ioutil.ReadFile(args[i])
			if err != nil {
				for k := range err.Error() {
					z01.PrintRune(rune(err.Error()[k]))
				}
				z01.PrintRune(10)
			}
			for i := range file {
				z01.PrintRune(rune(file[i]))
			}
		}

	}
}

func lenArgs(arr []string) int {
	l := 0
	for range arr {
		l++
	}
	return l
}
