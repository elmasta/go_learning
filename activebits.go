package piscine

func ActiveBits(n int) int {
	counter := 0
	for n != 0 {
		counter += n & 1
		n >>= 1
	}
	return counter
}
