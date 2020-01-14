package student

//Split a function that separates the words of a string and puts them in a string array.
func Split(str, charset string) []string {
	var count int
	if StrLen(str) < StrLen(charset) {
		array := []string{str}
		return array
	} else if str == charset && str == "" {
		array := []string{}
		return array
	} else if str != "" && charset == "" {
		count = StrLen(str) - 1
	} else {
		count = 0
		for i := 0; i <= StrLen(str)-StrLen(charset); {
			if str[i] == charset[0] {
				temp := str[i : i+StrLen(charset)]
				if temp == charset {
					count++
					if count+1 < 0 {
						array := make([]string, 0)
						return array
					}
					i = i + StrLen(charset)
					continue
				}
			}
			i++
		}
	}
	array := make([]string, count+1)
	tempStr := ""
	index := 0
	for i := 0; i < StrLen(str); {
		if charset == "" {
			array[index] = string(str[i])
			index++
			i++
			continue
		}
		if str[i] == charset[0] && (StrLen(str)-i) >= StrLen(charset) {
			temp := str[i : i+StrLen(charset)]
			if temp == charset {
				array[index] = tempStr
				index++
				tempStr = ""
				i = i + StrLen(charset)
				continue
			}
		}
		tempStr += string(rune(str[i]))
		i++
	}
	if tempStr != "" {
		array[index] = tempStr
	}
	return array
}
