package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	} else {
		toReturn := make([]int, max-min)
		iCount := 0
		for i := min; i < max; i++ {
			toReturn[iCount] = i
			iCount++
		}
		return toReturn
	}
}
