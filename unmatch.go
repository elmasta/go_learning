package piscine

func Unmatch(a []int) int {
	for _, v := range a {
		temp := 0
		for _, subV := range a {
			if v == subV {
				temp++
			}
		}
		if temp%2 != 0 {
			return v
		}
	}
	return -1
}
