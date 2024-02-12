package piscine

func IsUpper(s string) bool {
	sR := []rune(s)
	toReturn := true
	for _, v := range sR {
		if v < 65 || v > 90 {
			toReturn = false
			break
		}
	}
	return toReturn
}
