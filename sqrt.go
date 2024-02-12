package piscine

func Sqrt(nb int) int {
	result := 0
	if nb > 0 {
		for i := 0; i > -1; i++ {
			if i*i == nb {
				result = i
				break
			} else if i*i > nb {
				break
			}
		}
	}
	return result
}
