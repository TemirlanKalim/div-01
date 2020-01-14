package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if lnArgs(args) != 4 {
		return
	}
	errMess1 := "No division by 0"
	errMess2 := "No modulo by 0"
	errMess3 := "0"
	a, err := atoi2(args[1])
	b, err2 := atoi2(args[3])
	var c int64
	if err != "" || err2 != "" {
		printErrMess(errMess3)
		return
	}
	switch args[2] {
	case "+":
		c = a + b
		if a > 0 {
			if b > 0 {
				if c < a || c < b {
					printErrMess(errMess3)
					return
				}
			}
		} else if a < 0 {
			if b < 0 {
				if c > a || c > b {
					printErrMess(errMess3)
					return
				}
			}
		}
	case "-":
		c = a - b
		if a > 0 {
			if b < 0 {
				if c < a || c < b {
					printErrMess(errMess3)
					return
				}
			}
		} else if a < 0 {
			if b > 0 {
				if c > a || c > b {
					printErrMess(errMess3)
					return
				}
			}
		}
	case "*":
		c = a * b
		if a != 0 && b != 0 {
			if (c/b) != a || c/a != b {
				printErrMess(errMess3)
				return
			}
		}

	case "/":
		if b == 0 {
			printErrMess(errMess1)
			return
		}
		c = a / b
		if b == -1 {
			if (c*b) != a || (a/c) != b {
				printErrMess(errMess3)
				return
			}
		}
	case "%":
		if b == 0 {
			printErrMess(errMess2)
			return
		}
		c = a % b
	default:
		printErrMess(errMess3)
		return
	}
	printErrMess(itoa2(c))
}

func itoa2(n int64) string {
	str := ""
	if n == -9223372036854775808 {
		str = "-9223372036854775808"
	} else if n == 0 {
		str = "0"
	} else {
		minus := false
		if n < 0 {
			minus = true
			n *= -1
		}
		for n > 0 {
			if n < 10 {
				str = string(48+n) + str
				break
			} else {
				str = string(48+(n%10)) + str
				n /= 10
			}
		}
		if minus {
			str = string('-') + str
		}
	}
	return str
}
func printErrMess(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

func lnArgs(args []string) int {
	len := 0
	for range args {
		len++
	}
	return len
}

func atoi2(s string) (int64, string) {
	var minus bool
	var nbr int64
	if s == "-9223372036854775808" {
		nbr = -9223372036854775808
		return nbr, ""
	}
	len := 0
	for range s {
		len++
	}
	if len == 0 {
		return 0, "Error"
	}
	minus = false
	nbr = 0
	for i, v := range s {
		if v == '+' && i == 0 {
			continue
		} else if v == '+' && i != 0 {
			return 0, "Error"
		} else if v == '-' && i == 0 {
			minus = true
			continue
		} else if v == '-' && i != 0 {
			return 0, "Error"
		} else if v >= '0' && v <= '9' {
			var temp int64
			for k := '0'; k < v; k++ {
				temp++
			}
			temp2 := nbr
			nbr = nbr*10 + temp
			if temp2 > nbr {
				return 0, "Error"
			}
		} else {
			return 0, "Error"
		}
	}
	if minus {
		return nbr * -1, ""
	} else {
		return nbr, ""
	}
}
