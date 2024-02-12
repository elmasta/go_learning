package piscine

func Max(a []int) int {
	temp := 0
	for _, v := range a {
		if v > temp {
			temp = v
		}
	}
	return temp
}
