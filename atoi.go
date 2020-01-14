package student

//Atoi function
func Atoi(s string) int {
	nb := StrLen(s)
	newstr := []rune(s)
	s = ""
	final := 0
	num := 0
	ten := 10
	plus := 0
	minus := 0
	for i := 0; i < nb; i++ {
		if minus > 1 || plus > 1 || minus+plus > 1 {
			return 0
		}
		if newstr[i] >= '0' && newstr[i] <= '9' {
			num = 0
			for j := '0'; j < newstr[i]; j++ {
				num++
			}
			final = final*ten + num
		} else if newstr[i] == '+' && i == 0 {
			plus++
		} else if newstr[i] == '-' && i == 0 {
			minus++
		} else {
			return 0
		}
	}
	if minus == 1 {
		final = -final
	}
	return final
}
