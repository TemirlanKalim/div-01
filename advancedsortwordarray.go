package student

//AdvancedSortWordArr a function  that sorts a string array, based on the function f passed in parameter.
func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array)-1; i++ {
		for k := i + 1; k < len(array); k++ {
			if f(array[i], array[k]) > 0 {
				temp := array[i]
				array[i] = array[k]
				array[k] = temp
			}
		}
	}
}

//Compare of strings
func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	}
	return 1
}
