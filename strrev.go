package piscine

func StrRev(s string) string {
	stringR := ""
	for i := len(s) - 1; i >= 0; i-- {
		stringR = stringR + string(s[i])
	}
	return stringR
}
