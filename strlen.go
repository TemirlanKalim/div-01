package student

//StrLen length of string
func StrLen(str string) int {
	newstr := []rune(str)
	nb := 0
	for range newstr {
		nb++
	}
	return nb
}
