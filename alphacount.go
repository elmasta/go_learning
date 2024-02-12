package piscine

func AlphaCount(s string) int {
	count := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			count++
		} else if v >= 'A' && v <= 'Z' {
			count++
		}
	}
	return count
}
