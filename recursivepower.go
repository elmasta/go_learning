package piscine

func RecursivePower(nb int, power int) int {
	result := 1
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	} else if power < 0 {
		result = 0
	}
	return result
}
