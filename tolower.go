package piscine

func ToLower(s string) string {
	sR := []rune(s)
	for i, v := range sR {
		if v >= 'A' && v <= 'Z' {
			sR[i] = v + 32
		}
	}
	return string(sR)
}
