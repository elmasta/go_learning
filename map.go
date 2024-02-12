package piscine

func Map(f func(int) bool, a []int) []bool {
	var toReturn []bool
	for _, v := range a {
		toReturn = append(toReturn, f(v))
	}
	return toReturn
}
