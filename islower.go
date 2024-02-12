package piscine

func IsLower(s string) bool {
	sR := []rune(s)
	toReturn := true
	for _, v := range sR {
		if v < 97 || v > 122 {
			toReturn = false
			break
		}
	}
	return toReturn
}
