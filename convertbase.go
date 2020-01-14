package student

//ConvertBase a function that returns the conversion of a string number from one string baseFrom to one string baseTo.
func ConvertBase(nbr, baseFrom, baseTo string) string {
	if !checkbase(baseFrom) || !checkbase(baseTo) {
		str := "NV"
		return str
	}
	str := PrintNbrBase2(AtoiBase(nbr, baseFrom), baseTo)
	return str
}

//PrintNbrBase2 func
func PrintNbrBase2(nbr int, base string) string {
	str := ""
	minus := false
	if nbr < 0 {
		minus = true
	}
	if nbr == 0 {
		str = str + string(rune(base[0]))
		return str
	}
	for nbr != 0 {
		if nbr%StrLen(base) < 0 {
			str = string(base[(nbr%StrLen(base))*-1]) + str
		} else {
			str = string(base[(nbr%StrLen(base))]) + str
		}
		nbr = nbr / StrLen(base)
	}
	if minus {
		str = string('-') + str
	}
	return str
}
