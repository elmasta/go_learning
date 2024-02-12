package piscine

func IsNumeric(s string) bool {
	sR := []rune(s)
	toReturn := true
	for _, v := range sR {
		if v < 48 || v > 57 {
			toReturn = false
			break
		}
	}
	return toReturn
}
