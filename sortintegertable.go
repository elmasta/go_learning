package piscine

func SwapInTable(table []int, index int) []int {
	a, b := table[index], table[index+1]
	table[index], table[index+1] = b, a
	return table
}

func SortIntegerTable(table []int) {
	i := 0
	for true {
		if len(table) > 1 {
			if i == len(table)-2 {
				if table[i] > table[i+1] {
					table = SwapInTable(table, i)
					i = 0
				} else {
					break
				}
			} else if table[i] > table[i+1] {
				table = SwapInTable(table, i)
				i = 0
			} else {
				i++
			}
		} else {
			break
		}
	}
}
