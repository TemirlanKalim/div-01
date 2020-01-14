package student

//ActiveBits a function that returns the number of active bits (bits with the value 1) in the binary representation of an integer number.
func ActiveBits(n int) uint {
	var i uint
	if n == 0 {
		i = 0
	} else if n < 0 {
		for n != 0 {
			if n%2 == -1 {
				i++
				n /= 2
				continue
			}
			n /= 2
		}
		i++
	} else {
		for n != 0 {
			if n%2 == 1 {
				i++
				n /= 2
				continue
			}
			n /= 2
		}
	}
	return i
}
