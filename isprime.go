package piscine

func IsPrime(nb int) bool {
	isPrime := false
	if nb > 2 {
		for i := 2; i <= nb; i++ {
			if i == nb {
				isPrime = true
			}
			if nb%i == 0 {
				break
			}
		}
	} else if nb == 2 {
		isPrime = true
	}
	return isPrime
}
