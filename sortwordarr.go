package piscine

func SwapInTableString(table []string, index int) []string {
	a, b := table[index], table[index+1]
	table[index], table[index+1] = b, a
	return table
}

func SortWordArr(a []string) {
	i := 0
	for true {
		if len(a) > 1 {
			if i == len(a)-2 {
				if a[i] > a[i+1] {
					a = SwapInTableString(a, i)
					i = 0
				} else {
					break
				}
			} else if a[i] > a[i+1] {
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
