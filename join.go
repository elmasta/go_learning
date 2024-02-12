package piscine

func Join(strs []string, sep string) string {
	toReturn := ""
	for i, v := range strs {
		toReturn += v
		if i+1 != len(strs) {
			toReturn += sep
		}
	}
	return toReturn
}
