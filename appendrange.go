package piscine

func AppendRange(min, max int) []int {
	var toReturn []int
	if min >= max {
		return []int(nil)
	} else {
		for i := min; i < max; i++ {
			toReturn = append(toReturn, i)
		}
		return toReturn
	}
}
