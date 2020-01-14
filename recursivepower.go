package student

//RecursivePower function that returns the power of the int passed as parameter.
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return (nb * RecursivePower(nb, (power-1)))
	}
}
