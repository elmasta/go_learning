package piscine

func IsAlpha(s string) bool {
	sR := []rune(s)
	toReturn := true
	for _, v := range sR {
		if v >= 97 && v <= 122 {
			toReturn = true
		} else if v >= 65 && v <= 90 {
			toReturn = true
		} else if v >= 48 && v <= 57 {
			toReturn = true
		} else {
			toReturn = false
			break
		}
	}
	return toReturn
}
