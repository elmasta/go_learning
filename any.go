package piscine

func Any(f func(string) bool, a []string) bool {
	toReturn := false
	for _, v := range a {
		if f(v) {
			toReturn = true
			break
		}
	}
	return toReturn
}
