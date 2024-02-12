package piscine

func Compact(ptr *[]string) int {
	var list []string
	counter := 0
	for _, v := range *ptr {
		if len(v) > 0 {
			list = append(list, v)
			counter++
		}
	}
	*ptr = list
	return counter
}
