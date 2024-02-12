package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	i := 0
	for true {
		if len(a) > 1 {
			if i == len(a)-2 {
				if f(a[i], a[i+1]) == 1 {
					a = SwapInTableString(a, i)
					i = 0
				} else {
					break
				}
			} else if f(a[i], a[i+1]) == 1 {
				a = SwapInTableString(a, i)
				i = 0
			} else {
				i++
			}
		} else {
			break
		}
	}
}
