package student

//AtoiBase a function that takes a string number and its string base in parameters and returns its conversion as an int.
func AtoiBase(s string, base string) int {
	if !checkbase(base) || !checkString(s, base) {
		return 0
	}
	result := 0
	s = strRev(s)
	for i := 0; i < lenStr2(s); i++ {
		for k := range base {
			if s[i] == base[k] {
				result = (k * rPower(lenStr2(base), i)) + result
			}
		}
	}
	return result
}

func checkbase(base string) bool {
	if lenStr2(base) < 2 {
		return false
	}
	for _, v := range base {
		if v == 43 || v == 45 || v < 33 || v > 125 {
			return false
		}
	}
	for i := 0; i < lenStr2(base)-1; i++ {
		for k := i + 1; k < lenStr2(base); k++ {
			if base[i] == base[k] {
				return false
			}
		}
	}
	return true
}

func checkString(str string, base string) bool {
	if lenStr2(str) < 1 {
		return false
	}
	for _, v := range str {
		if v == 43 || v == 45 || v < 33 || v > 125 {
			return false
		}
		for i := 0; i < lenStr2(base); i++ {
			if v == rune(base[i]) {
				break
			} else if v != rune(base[i]) && i < lenStr2(base)-1 {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func lenStr2(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func rPower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		nb *= RecursivePower(nb, power-1)
	}
	return nb
}

func strRev(s string) string {
	str := ""
	for _, v := range s {
		str = string(v) + str
	}
	return str
}
