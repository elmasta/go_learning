package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 1 && nb < 26 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	} else if nb == 0 || nb == 1 {
		result = 1
	} else {
		result = 0
	}
	return result
}
