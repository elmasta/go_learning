package piscine

func SplitWhiteSpaces(s string) []string {
	var toReturn []string
	start := -1
	for i, v := range s {
		if (v == ' ' || v == '\t' || v == '\n') && start == -1 {
			start = -1
		} else if start == -1 {
			start = i
		}
		if (v == ' ' || v == '\t' || v == '\n') && start > -1 {
			toReturn = append(toReturn, s[start:i])
			start = -1
		} else if i+1 == len(s) && start > -1 {
			toReturn = append(toReturn, s[start:i+1])
		}
	}
	return toReturn
}
