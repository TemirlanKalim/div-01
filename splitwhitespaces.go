package student

//SplitWhiteSpaces a function that separates the words of a string and puts them in a string array.
func SplitWhiteSpaces(str string) []string {
	for str[0] == 9 || str[0] == 32 || str[0] == 10 {
		str = str[1:]
	}
	for str[lenStr2(str)-1] == 9 || str[lenStr2(str)-1] == 32 || str[lenStr2(str)-1] == 10 {
		str = str[:lenStr2(str)-1]
	}
	count := 0
	for i := range str {
		if str[i] == 9 || str[i] == 32 || str[i] == 10 {
			if str[i-1] != 9 && str[i-1] != 32 && str[i-1] != 10 {
				count++
				if count+1 < 0 {
					array := make([]string, 0)
					return array
				}
			}
		}
	}
	array := make([]string, count+1)
	index := 0
	tempStr := ""
	for i := range str {
		if str[i] == 9 || str[i] == 32 || str[i] == 10 {
			if tempStr != "" {
				array[index] = tempStr
				tempStr = ""
				index++
			} else {
				continue
			}
		} else {
			tempStr = tempStr + string(rune(str[i]))
		}
	}
	array[index] = tempStr
	return array
}
