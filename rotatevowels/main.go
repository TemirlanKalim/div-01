package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if lenArgs(args) == 1 {
		z01.PrintRune(10)
		return
	}
	args = args[1:]
	arrVow := []rune{65, 69, 73, 79, 85, 97, 101, 105, 111, 117}
	strOfVow := ""
	ind := 0
	for i := range args {
		for k := range args[i] {
			for j := range arrVow {
				if arrVow[j] == rune(args[i][k]) {
					strOfVow = string(args[i][k]) + strOfVow
					break
				}
			}
		}
	}

	for i := range args {
		for k := range args[i] {
			for j := range arrVow {
				if arrVow[j] == rune(args[i][k]) {
					z01.PrintRune(rune(strOfVow[ind]))
					ind++
					break
				} else if arrVow[j] != rune(args[i][k]) && j != 9 {
					continue
				} else {
					z01.PrintRune(rune(args[i][k]))
					break
				}
			}
		}
		if i != lenArgs(args)-1 {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune(10)
}
func lenArgs(arr []string) int {
	l := 0
	for range arr {
		l++
	}
	return l
}
