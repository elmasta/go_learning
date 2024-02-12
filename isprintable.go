package piscine

func IsPrintable(s string) bool {
	sR := []rune(s)
	toReturn := true
	for _, v := range sR {
		if v > 31 && v < 127 {
			toReturn = true
		} else {
			toReturn = false
			break
		}
	}
	return toReturn
}
