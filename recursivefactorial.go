package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb > 1 && nb < 26 {
		return nb * RecursiveFactorial(nb-1)
	} else if nb == 0 || nb == 1 {
		result = 1
	} else {
		result = 0
	}
	return result
}
