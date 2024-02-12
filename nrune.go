package piscine

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		iRune := []rune(s)
		newRune := iRune[n-1]
		return newRune
	} else {
		return 0
	}
}
