package piscine

func ToUpper(s string) string {
	sR := []rune(s)
	for i, v := range sR {
		if v >= 'a' && v <= 'z' {
			sR[i] = v - 32
		}
	}
	return string(sR)
}
