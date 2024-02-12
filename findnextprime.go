package piscine

func FindNextPrime(nb int) int {
	nextPrime := 2
	if nb > 2 && nb != 3 {
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				nb++
				i = 1
			}
		}
		nextPrime = nb
	} else if nb == 3 {
		nextPrime = 3
	}
	return nextPrime
}
