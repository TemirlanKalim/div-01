package student

import (
	"github.com/01-edu/z01"
)

//PrintNbrBase function that prints an int in a string base passed in parameters.
func PrintNbrBase(nbr int, base string) {
	len := StrLen(base)
	runes := []rune(base)
	valid := true
	if len < 2 {
		valid = false
	} else {
		for i := 0; i < len; i++ {
			if runes[i] == '-' || runes[i] == '+' {
				valid = false
			}
			for j := i + 1; j < len; j++ {
				if runes[i] == runes[j] {
					valid = false
				}
			}
		}
	}
	if !valid {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr == 0 {
			z01.PrintRune(runes[0])
		} else {
			if nbr < 0 {
				z01.PrintRune('-')
			}
			PrintNbrBaseRec(nbr, runes, len)
		}
	}
}

//PrintNbrBaseRec function to convert number
func PrintNbrBaseRec(n int, runes []rune, len int) {
	if n/len != 0 {
		PrintNbrBaseRec(n/len, runes, len)
	}
	mod := n % len
	if mod < 0 {
		mod = -mod
	}
	z01.PrintRune(runes[mod])
}
