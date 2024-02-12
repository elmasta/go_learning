package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
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
	return table[2]
}
