package piscine

func Rot14(s string) string {
	toReturn := ""
	for _, v := range s {
		if rune(v) >= 'A' && rune(v) <= 'Z' {
			v += 14
			if v > 'Z' {
				v = v - 26
			}
			toReturn += string(v)
		} else if rune(v) >= 'a' && rune(v) <= 'z' {
			v += 14
			if v > 'z' {
				v = v - 26
			}
			toReturn += string(v)
		} else {
			toReturn += string(v)
		}
	}
	return toReturn
}
